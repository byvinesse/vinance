//
//  HomeView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/05/26.
//
//  Reference design: design/mock/mock-v1.html (VNC-38)
//
//  Restructured home screen built on the Vinance Mobile App Design system.
//  Sections, top to bottom:
//    • Header — two-line Sora greeting · 3-state privacy mask cycle · avatar
//    • Balance card — TOTAL BALANCE + mask pill, hero amount with split
//      decimal, monthly delta, INCOME / EXPENSES / SAVED pills
//    • Quick actions — Expense · Income · Transfer · Report
//    • Accounts — horizontal strip of cards + dashed "Add account" tile
//    • Recent — grouped transaction list (Today / Yesterday / weekday)
//    • Live currencies — 2-col FX grid against the user's base currency
//
//  Privacy mask states (matching the reference):
//    .visible       → all amounts shown, green pill "Visible"
//    .amountHidden  → amounts blurred, percentages still visible, yellow pill
//    .allHidden     → amounts AND percentages blurred, red pill
//

import SwiftUI

// MARK: - Privacy Mask ───────────────────────────────────────────────────────

enum MaskLevel: Int, CaseIterable {
    case visible       = 0
    case amountHidden  = 1
    case allHidden     = 2

    var next: MaskLevel {
        MaskLevel(rawValue: (rawValue + 1) % MaskLevel.allCases.count) ?? .visible
    }

    var label: String {
        switch self {
        case .visible:      return "Visible"
        case .amountHidden: return "฿ Hidden"
        case .allHidden:    return "All Hidden"
        }
    }

    var pillColor: Color {
        switch self {
        case .visible:      return .vGreen
        case .amountHidden: return .vYellow
        case .allHidden:    return .vRed
        }
    }

    /// Should hero amounts be blurred at this level?
    var hidesAmount: Bool { self != .visible }

    /// Should percentages / deltas be blurred at this level?
    var hidesPercent: Bool { self == .allHidden }
}

private extension View {
    /// Privacy-mask blur, layout-stable.
    @ViewBuilder
    func privacyMasked(_ hidden: Bool) -> some View {
        if hidden {
            self.blur(radius: 7).opacity(0.35).allowsHitTesting(false)
        } else {
            self
        }
    }
}

// MARK: - Local Model Types ──────────────────────────────────────────────────

/// Rendering currency for an account / amount.
struct HomeCurrency: Equatable {
    let code: String   // "THB", "IDR", "USD", "EUR"
    let symbol: String // "฿", "Rp", "$", "€"
    let decimals: Int

    static let thb = HomeCurrency(code: "THB", symbol: "฿",  decimals: 2)
    static let idr = HomeCurrency(code: "IDR", symbol: "Rp", decimals: 0)
    static let usd = HomeCurrency(code: "USD", symbol: "$",  decimals: 2)
    static let eur = HomeCurrency(code: "EUR", symbol: "€",  decimals: 2)
    static let sgd = HomeCurrency(code: "SGD", symbol: "S$", decimals: 2)
}

/// EU-style number formatting: `1.234.567,89`.
/// Mirrors the reference design's `fmt()` helper.
struct AmountParts: Equatable {
    let symbol: String
    let integer: String
    let decimal: String   // empty when currency.decimals == 0
}

enum MoneyFormatter {
    static func split(_ value: Double, in ccy: HomeCurrency) -> AmountParts {
        let abs = Swift.abs(value)
        let f = NumberFormatter()
        f.locale = Locale(identifier: "de_DE")          // 1.234,56
        f.numberStyle = .decimal
        f.minimumFractionDigits = ccy.decimals
        f.maximumFractionDigits = ccy.decimals
        f.usesGroupingSeparator = true
        let s = f.string(from: NSNumber(value: abs)) ?? "\(abs)"
        if ccy.decimals == 0 {
            return AmountParts(symbol: ccy.symbol, integer: s, decimal: "")
        }
        let parts = s.split(separator: ",", maxSplits: 1, omittingEmptySubsequences: false)
        let intPart = parts.first.map(String.init) ?? s
        let decPart = parts.count > 1 ? String(parts[1]) : ""
        return AmountParts(symbol: ccy.symbol, integer: intPart, decimal: decPart)
    }

    /// Single-line representation, optionally signed. Used by transactions / pills.
    static func formatted(_ value: Double, in ccy: HomeCurrency, sign: Sign = .none) -> String {
        let p = split(value, in: ccy)
        let body = p.decimal.isEmpty ? "\(p.symbol)\(p.integer)" : "\(p.symbol)\(p.integer),\(p.decimal)"
        switch sign {
        case .plus:  return "+\(body)"
        case .minus: return "−\(body)"   // U+2212 minus
        case .none:  return body
        }
    }

    enum Sign { case none, plus, minus }
}

/// Account model used on the home screen.
struct HomeAccount: Identifiable {
    let id = UUID()
    let bank: String
    let accountName: String
    let amount: Double
    let currency: HomeCurrency
    let gradient: [Color]
    let dotColor: Color
}

/// Transaction display type.
enum HomeTxKind { case expense, income, transfer }

/// Transaction model used on the home screen.
struct HomeTxItem: Identifiable {
    let id = UUID()
    let kind: HomeTxKind
    let categoryName: String
    let categoryColor: Color
    let emoji: String
    let account: String
    let toAccount: String?            // populated for transfers
    let description: String
    let amount: Double
    let currency: HomeCurrency
    let time: String
    let dateKey: String               // grouping key, e.g. "today", "mon-5"

    var sign: MoneyFormatter.Sign {
        switch kind {
        case .expense:  return .minus
        case .income:   return .plus
        case .transfer: return .none
        }
    }

    var amountColor: Color {
        switch kind {
        case .expense:  return .vRed
        case .income:   return .vGreen
        case .transfer: return .vText
        }
    }

    var formattedAmount: String {
        MoneyFormatter.formatted(amount, in: currency, sign: sign)
    }
}

/// FX pair shown in the Live currencies grid.
struct FxPair: Identifiable {
    let id = UUID()
    let from: String     // e.g. "USD"
    let to: String       // e.g. "THB"
    let rate: Double     // 1 [from] = N [to]
    let deltaPct: Double // signed % change vs yesterday

    var isUp: Bool { deltaPct >= 0 }
}

// MARK: - Date Group Labels ──────────────────────────────────────────────────

private enum DateGroup {
    static let labels: [String: String] = [
        "today":     "Today",
        "yesterday": "Yesterday",
        "mon-5":     "Mon, 5 May",
        "sun-4":     "Sun, 4 May",
        "sat-3":     "Sat, 3 May",
        "fri-2":     "Fri, 2 May"
    ]
    static func label(for key: String) -> String { labels[key] ?? key }
}

// MARK: - Mock Data ──────────────────────────────────────────────────────────

enum MockHomeData {

    static let totalBalance: Double = 284_520
    static let monthIncome:  Double = 68_000
    static let monthExpense: Double = 34_280
    static let monthSaved:   Double = 33_720
    static let monthDelta:   Double = 4.2     // % vs last month

    static let accounts: [HomeAccount] = [
        HomeAccount(
            bank: "Kasikorn Bank",
            accountName: "Savings Account",
            amount: 142_300,
            currency: .thb,
            gradient: [Color(hex: 0x1e1b4b), Color(hex: 0x312e81)],
            dotColor: Color(hex: 0xf59e0b)
        ),
        HomeAccount(
            bank: "Bangkok Bank",
            accountName: "Current Account",
            amount: 89_220,
            currency: .thb,
            gradient: [Color(hex: 0x064e3b), Color(hex: 0x065f46)],
            dotColor: Color(hex: 0x34d399)
        ),
        HomeAccount(
            bank: "BCA Mandiri",
            accountName: "Credit Card · IDR",
            amount: 12_500_000,
            currency: .idr,
            gradient: [Color(hex: 0x1e3a5f), Color(hex: 0x1e3a8a)],
            dotColor: Color(hex: 0x60a5fa)
        )
    ]

    static let transactions: [HomeTxItem] = [
        HomeTxItem(kind: .expense,  categoryName: "Groceries",     categoryColor: .vYellow, emoji: "🛒",
                   account: "Kasikorn Bank", toAccount: nil, description: "Tops Supermarket",
                   amount: 840,    currency: .thb, time: "11:24", dateKey: "today"),
        HomeTxItem(kind: .expense,  categoryName: "Food & Drink",  categoryColor: .vBlue,   emoji: "☕",
                   account: "Kasikorn Bank", toAccount: nil, description: "Starbucks Siam",
                   amount: 180,    currency: .thb, time: "08:50", dateKey: "today"),
        HomeTxItem(kind: .income,   categoryName: "Income",        categoryColor: .vGreen,  emoji: "💼",
                   account: "Bangkok Bank",  toAccount: nil, description: "Monthly Salary",
                   amount: 68_000, currency: .thb, time: "09:00", dateKey: "yesterday"),
        HomeTxItem(kind: .expense,  categoryName: "Subscriptions", categoryColor: .vPink,   emoji: "🎵",
                   account: "BCA Mandiri",   toAccount: nil, description: "Spotify Premium",
                   amount: 99,     currency: .thb, time: "00:00", dateKey: "yesterday"),
        HomeTxItem(kind: .expense,  categoryName: "Food",          categoryColor: .vOrange, emoji: "🍜",
                   account: "Kasikorn Bank", toAccount: nil, description: "Som Tam Jay So",
                   amount: 220,    currency: .thb, time: "13:10", dateKey: "mon-5")
    ]

    static let fxPairs: [FxPair] = [
        FxPair(from: "USD", to: "THB", rate: 34.5,  deltaPct:  0.12),
        FxPair(from: "SGD", to: "THB", rate: 25.6,  deltaPct:  0.08),
        FxPair(from: "THB", to: "THB", rate: 1.0,   deltaPct:  0.06),
        FxPair(from: "EUR", to: "THB", rate: 37.4,  deltaPct:  0.21)
    ]
}

// MARK: - HomeView ───────────────────────────────────────────────────────────

struct HomeView: View {

    @State private var maskLevel: MaskLevel = .visible

    var body: some View {
        ZStack {
            Color.vBg.ignoresSafeArea()

            ScrollView(showsIndicators: false) {
                VStack(spacing: 0) {
                    HomeHeaderBar(
                        maskLevel: maskLevel,
                        onCycleMask: cycleMask
                    )

                    BalanceCard(
                        total:     MockHomeData.totalBalance,
                        income:    MockHomeData.monthIncome,
                        expenses:  MockHomeData.monthExpense,
                        saved:     MockHomeData.monthSaved,
                        deltaPct:  MockHomeData.monthDelta,
                        maskLevel: maskLevel
                    )
                    .padding(.horizontal, 24)
                    .padding(.top, 14)

                    SectionHeader(title: "Quick actions")
                    QuickActionsRow()
                        .padding(.horizontal, 24)
                        .padding(.top, 12)

                    SectionHeader(title: "Accounts", trailing: "Manage")
                    AccountsStrip(
                        accounts: MockHomeData.accounts,
                        maskLevel: maskLevel
                    )

                    SectionHeader(title: "Recent", trailing: "See all")
                    RecentTransactionsList(
                        transactions: MockHomeData.transactions,
                        maskLevel: maskLevel
                    )
                    .padding(.top, 4)

                    SectionHeader(title: "Live currencies", trailing: "Manage")
                    LiveCurrenciesGrid(
                        pairs: MockHomeData.fxPairs,
                        maskLevel: maskLevel
                    )
                    .padding(.top, 4)

                    Color.clear.frame(height: 96) // breathing room above the tab bar
                }
            }
        }
        .navigationBarHidden(true)
    }

    private func cycleMask() {
        withAnimation(.easeInOut(duration: 0.22)) {
            maskLevel = maskLevel.next
        }
    }
}

// MARK: - Header ─────────────────────────────────────────────────────────────

private struct HomeHeaderBar: View {
    let maskLevel: MaskLevel
    let onCycleMask: () -> Void

    var body: some View {
        HStack(alignment: .center, spacing: 10) {
            VStack(alignment: .leading, spacing: 0) {
                Text("Good morning,")
                    .font(.system(size: 15, weight: .medium))
                    .foregroundColor(.vText2)
                Text("Alex 👋")
                    .font(.system(size: 26, weight: .bold, design: .rounded))
                    .foregroundColor(.vText)
                Text("\(headerDate) · Bangkok")
                    .font(.system(size: 12))
                    .foregroundColor(.vText3)
                    .padding(.top, 4)
            }
            Spacer()

            Button(action: onCycleMask) {
                ZStack {
                    Circle()
                        .fill(Color.white.opacity(0.08))
                        .overlay(Circle().stroke(Color.white.opacity(0.14), lineWidth: 1))
                    Image(systemName: maskLevel.eyeSFSymbol)
                        .font(.system(size: 13, weight: .medium))
                        .foregroundColor(maskLevel.eyeIconColor)
                }
                .frame(width: 36, height: 36)
            }
            .accessibilityLabel("Privacy mask, currently \(maskLevel.label)")

            ZStack {
                Circle()
                    .fill(LinearGradient(colors: [.vAccent, .vPink],
                                         startPoint: .topLeading,
                                         endPoint: .bottomTrailing))
                Text("A")
                    .font(.system(size: 14, weight: .bold))
                    .foregroundColor(.white)
            }
            .frame(width: 36, height: 36)
        }
        .padding(.horizontal, 24)
        .padding(.top, 16)
        .padding(.bottom, 4)
    }

    private var headerDate: String {
        let f = DateFormatter()
        f.dateFormat = "EEEE, d MMM"
        return f.string(from: Date())
    }
}

private extension MaskLevel {
    var eyeSFSymbol: String {
        switch self {
        case .visible:      return "eye"
        case .amountHidden: return "eye.fill"
        case .allHidden:    return "eye.slash.fill"
        }
    }

    var eyeIconColor: Color {
        switch self {
        case .visible:      return .vText2
        case .amountHidden: return .vYellow
        case .allHidden:    return .vRed
        }
    }
}

// MARK: - Balance Card ───────────────────────────────────────────────────────

private struct BalanceCard: View {
    let total:     Double
    let income:    Double
    let expenses:  Double
    let saved:     Double
    let deltaPct:  Double
    let maskLevel: MaskLevel

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {

            HStack(alignment: .top) {
                Text("TOTAL BALANCE")
                    .font(.system(size: 10, weight: .semibold))
                    .kerning(0.9)
                    .foregroundColor(.vText3)
                Spacer()
                MaskPill(level: maskLevel)
            }

            HeroAmount(value: total, currency: .thb, maskLevel: maskLevel)
                .padding(.top, 6)

            HStack(spacing: 4) {
                Image(systemName: "arrowtriangle.up.fill").font(.system(size: 9))
                Text("+\(deltaText)% from last month")
                    .font(.system(size: 12))
            }
            .foregroundColor(.vGreen)
            .padding(.top, 6)
            .privacyMasked(maskLevel.hidesPercent)

            HStack(spacing: 10) {
                BalancePill(label: "INCOME",   value: income,   sign: .plus,  color: .vGreen, maskLevel: maskLevel)
                BalancePill(label: "EXPENSES", value: expenses, sign: .minus, color: .vRed,   maskLevel: maskLevel)
                BalancePill(label: "SAVED",    value: saved,    sign: .none,  color: .vBlue,  maskLevel: maskLevel)
            }
            .padding(.top, 18)
        }
        .padding(22)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(
            ZStack {
                LinearGradient(
                    colors: [Color(hex: 0x1a1730), Color(hex: 0x0f1128)],
                    startPoint: .topLeading,
                    endPoint: .bottomTrailing
                )
                RadialGradient(colors: [Color.vAccent.opacity(0.22), .clear],
                               center: .topTrailing,
                               startRadius: 0, endRadius: 180)
                RadialGradient(colors: [Color.vPink.opacity(0.10), .clear],
                               center: .bottomLeading,
                               startRadius: 0, endRadius: 110)
            }
        )
        .clipShape(RoundedRectangle(cornerRadius: 24))
        .overlay(
            RoundedRectangle(cornerRadius: 24)
                .stroke(Color.vAccent.opacity(0.18), lineWidth: 1)
        )
    }

    private var deltaText: String {
        // de_DE locale → uses comma decimal: "4,2"
        let f = NumberFormatter()
        f.locale = Locale(identifier: "de_DE")
        f.minimumFractionDigits = 1
        f.maximumFractionDigits = 1
        return f.string(from: NSNumber(value: deltaPct)) ?? "\(deltaPct)"
    }
}

/// Hero amount: large integer + smaller `,dec` tail. Layout-stable when masked.
private struct HeroAmount: View {
    let value: Double
    let currency: HomeCurrency
    let maskLevel: MaskLevel

    var body: some View {
        let p = MoneyFormatter.split(value, in: currency)
        HStack(alignment: .lastTextBaseline, spacing: 4) {
            Text(p.symbol)
                .font(.system(size: 22, weight: .medium, design: .rounded))
                .foregroundColor(.vText2)
            Text(p.integer)
                .font(.system(size: 44, weight: .bold, design: .rounded))
                .kerning(-0.8)
                .foregroundColor(.vText)
                .privacyMasked(maskLevel.hidesAmount)
                .lineLimit(1)
                .minimumScaleFactor(0.7)
            if !p.decimal.isEmpty {
                Text(",\(p.decimal)")
                    .font(.system(size: 22, weight: .medium, design: .rounded))
                    .foregroundColor(.vText2)
                    .privacyMasked(maskLevel.hidesAmount)
            }
        }
    }
}

private struct MaskPill: View {
    let level: MaskLevel
    var body: some View {
        HStack(spacing: 6) {
            Circle()
                .fill(level.pillColor)
                .frame(width: 5, height: 5)
                .shadow(color: level.pillColor, radius: 3)
            Text(level.label)
                .font(.system(size: 9, weight: .semibold))
                .kerning(0.9)
        }
        .padding(.horizontal, 10)
        .padding(.vertical, 5)
        .foregroundColor(level.pillColor)
        .background(level.pillColor.opacity(0.12))
        .clipShape(Capsule())
        .overlay(Capsule().stroke(level.pillColor.opacity(0.30), lineWidth: 1))
    }
}

private struct BalancePill: View {
    let label: String
    let value: Double
    let sign: MoneyFormatter.Sign
    let color: Color
    let maskLevel: MaskLevel

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(label)
                .font(.system(size: 9, weight: .semibold))
                .kerning(0.9)
                .foregroundColor(.vText3)
            Text(MoneyFormatter.formatted(value, in: .thb, sign: sign))
                .font(.system(size: 13, weight: .semibold, design: .rounded))
                .foregroundColor(color)
                .lineLimit(1)
                .minimumScaleFactor(0.6)
                .privacyMasked(maskLevel.hidesAmount)
        }
        .padding(.horizontal, 12)
        .padding(.vertical, 11)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(Color.white.opacity(0.04))
        .clipShape(RoundedRectangle(cornerRadius: 14))
        .overlay(
            RoundedRectangle(cornerRadius: 14)
                .stroke(Color.white.opacity(0.07), lineWidth: 1)
        )
    }
}

// MARK: - Section Header ─────────────────────────────────────────────────────

private struct SectionHeader: View {
    let title: String
    var trailing: String? = nil

    var body: some View {
        HStack(alignment: .firstTextBaseline) {
            Text(title)
                .font(.system(size: 16, weight: .semibold, design: .rounded))
                .foregroundColor(.vText)
            Spacer()
            if let trailing {
                Text(trailing)
                    .font(.system(size: 13))
                    .foregroundColor(.vAccent2)
            }
        }
        .padding(.horizontal, 24)
        .padding(.top, 26)
    }
}

// MARK: - Quick Actions ──────────────────────────────────────────────────────

private struct QuickActionsRow: View {
    private struct Action: Identifiable {
        let id = UUID()
        let emoji: String
        let label: String
    }
    private let actions: [Action] = [
        .init(emoji: "💸", label: "Expense"),
        .init(emoji: "💰", label: "Income"),
        .init(emoji: "🔄", label: "Transfer"),
        .init(emoji: "📊", label: "Report")
    ]

    var body: some View {
        HStack(spacing: 10) {
            ForEach(actions) { a in
                VStack(spacing: 8) {
                    ZStack {
                        RoundedRectangle(cornerRadius: 13)
                            .fill(Color.white.opacity(0.08))
                            .overlay(
                                RoundedRectangle(cornerRadius: 13)
                                    .stroke(Color.white.opacity(0.14), lineWidth: 1)
                            )
                        Text(a.emoji).font(.system(size: 20))
                    }
                    .frame(width: 42, height: 42)
                    Text(a.label)
                        .font(.system(size: 11, weight: .medium))
                        .foregroundColor(.vText2)
                }
                .frame(maxWidth: .infinity)
                .padding(.vertical, 14)
                .background(Color.vSurface)
                .clipShape(RoundedRectangle(cornerRadius: 18))
                .overlay(
                    RoundedRectangle(cornerRadius: 18)
                        .stroke(Color.white.opacity(0.07), lineWidth: 1)
                )
            }
        }
    }
}

// MARK: - Accounts Strip ─────────────────────────────────────────────────────

private struct AccountsStrip: View {
    let accounts: [HomeAccount]
    let maskLevel: MaskLevel

    var body: some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: 14) {
                ForEach(accounts) { AccountCard(account: $0, maskLevel: maskLevel) }
                AddAccountTile()
            }
            .padding(.horizontal, 24)
            .padding(.top, 10)
            .padding(.bottom, 4)
        }
    }
}

private struct AccountCard: View {
    let account: HomeAccount
    let maskLevel: MaskLevel

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack {
                Spacer()
                HStack(spacing: 4) {
                    Circle().fill(account.dotColor).frame(width: 18, height: 18)
                    Circle().fill(account.dotColor.opacity(0.4)).frame(width: 18, height: 18)
                }
            }
            .padding(.bottom, 14)

            Text(account.bank.uppercased())
                .font(.system(size: 10, weight: .semibold))
                .kerning(0.7)
                .foregroundColor(.white.opacity(0.70))

            Text(MoneyFormatter.formatted(account.amount, in: account.currency))
                .font(.system(size: 21, weight: .bold, design: .rounded))
                .kerning(-0.2)
                .foregroundColor(.white)
                .padding(.top, 8)
                .lineLimit(1)
                .minimumScaleFactor(0.75)
                .privacyMasked(maskLevel.hidesAmount)

            Text(account.accountName)
                .font(.system(size: 12))
                .foregroundColor(.white.opacity(0.60))
                .padding(.top, 2)
        }
        .padding(18)
        .frame(width: 196)
        .background(
            ZStack {
                LinearGradient(colors: account.gradient,
                               startPoint: .topLeading,
                               endPoint: .bottomTrailing)
                RadialGradient(colors: [Color.white.opacity(0.10), .clear],
                               center: .topTrailing,
                               startRadius: 0, endRadius: 90)
            }
        )
        .clipShape(RoundedRectangle(cornerRadius: 20))
        .overlay(
            RoundedRectangle(cornerRadius: 20)
                .stroke(Color.white.opacity(0.07), lineWidth: 1)
        )
    }
}

private struct AddAccountTile: View {
    var body: some View {
        VStack(spacing: 8) {
            ZStack {
                Circle().fill(Color.vAccent)
                Image(systemName: "plus")
                    .font(.system(size: 18, weight: .light))
                    .foregroundColor(.white)
            }
            .frame(width: 36, height: 36)

            Text("Add account")
                .font(.system(size: 12, weight: .medium))
                .foregroundColor(.vText2)
        }
        .frame(width: 196, height: 132)
        .background(Color.white.opacity(0.04))
        .clipShape(RoundedRectangle(cornerRadius: 20))
        .overlay(
            RoundedRectangle(cornerRadius: 20)
                .strokeBorder(style: StrokeStyle(lineWidth: 1, dash: [4, 3]))
                .foregroundColor(Color.white.opacity(0.14))
        )
    }
}

// MARK: - Recent Transactions ────────────────────────────────────────────────

private struct RecentTransactionsList: View {
    let transactions: [HomeTxItem]
    let maskLevel: MaskLevel

    private var grouped: [(String, [HomeTxItem])] {
        var index: [String: Int] = [:]
        var result: [(String, [HomeTxItem])] = []
        for tx in transactions {
            if let i = index[tx.dateKey] {
                result[i].1.append(tx)
            } else {
                index[tx.dateKey] = result.count
                result.append((tx.dateKey, [tx]))
            }
        }
        return result
    }

    var body: some View {
        VStack(spacing: 0) {
            ForEach(grouped, id: \.0) { key, items in
                Text(DateGroup.label(for: key))
                    .font(.system(size: 10, weight: .semibold))
                    .kerning(0.9)
                    .foregroundColor(.vText3)
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .padding(.horizontal, 24)
                    .padding(.top, 14)
                    .padding(.bottom, 4)
                ForEach(items) { tx in
                    TxRow(tx: tx, maskLevel: maskLevel)
                    if tx.id != items.last?.id {
                        Divider().background(Color.white.opacity(0.07))
                            .padding(.leading, 80)
                    }
                }
            }
        }
    }
}

private struct TxRow: View {
    let tx: HomeTxItem
    let maskLevel: MaskLevel

    var body: some View {
        HStack(spacing: 12) {
            ZStack {
                RoundedRectangle(cornerRadius: 14)
                    .fill(Color.white.opacity(0.08))
                    .overlay(
                        LinearGradient(colors: [tx.categoryColor.opacity(0.16),
                                                tx.categoryColor.opacity(0.04)],
                                       startPoint: .topLeading,
                                       endPoint: .bottomTrailing)
                            .clipShape(RoundedRectangle(cornerRadius: 14))
                    )
                    .overlay(
                        RoundedRectangle(cornerRadius: 14)
                            .stroke(Color.white.opacity(0.14), lineWidth: 1)
                    )
                Text(tx.emoji).font(.system(size: 20))
            }
            .frame(width: 44, height: 44)

            VStack(alignment: .leading, spacing: 2) {
                Text(tx.categoryName)
                    .font(.system(size: 13, weight: .semibold))
                    .foregroundColor(.vText)
                    .lineLimit(1)
                Text(tx.account + (tx.toAccount.map { " → \($0)" } ?? ""))
                    .font(.system(size: 11))
                    .foregroundColor(.vText3)
                    .lineLimit(1)
                Text(tx.description)
                    .font(.system(size: 11).italic())
                    .foregroundColor(.vText2)
                    .lineLimit(1)
            }

            Spacer(minLength: 8)

            VStack(alignment: .trailing, spacing: 3) {
                Text(tx.formattedAmount)
                    .font(.system(size: 14, weight: .semibold, design: .rounded))
                    .foregroundColor(tx.amountColor)
                    .privacyMasked(maskLevel.hidesAmount)
                Text(tx.time)
                    .font(.system(size: 10))
                    .foregroundColor(.vText3)
            }
        }
        .padding(.horizontal, 24)
        .padding(.vertical, 12)
    }
}

// MARK: - Live Currencies Grid ───────────────────────────────────────────────

private struct LiveCurrenciesGrid: View {
    let pairs: [FxPair]
    let maskLevel: MaskLevel

    private let columns = [
        GridItem(.flexible(), spacing: 8),
        GridItem(.flexible(), spacing: 8)
    ]

    var body: some View {
        LazyVGrid(columns: columns, spacing: 8) {
            ForEach(pairs) { FxCard(pair: $0, maskLevel: maskLevel) }
        }
        .padding(.horizontal, 24)
    }
}

private struct FxCard: View {
    let pair: FxPair
    let maskLevel: MaskLevel

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            HStack(spacing: 4) {
                Text(pair.from)
                    .foregroundColor(.vText)
                Text("/")
                    .foregroundColor(.vText3.opacity(0.6))
                Text(pair.to)
                    .foregroundColor(.vText3)
            }
            .font(.system(size: 10, weight: .semibold))
            .kerning(0.8)

            Text(formatRate(pair.rate))
                .font(.system(size: 18, weight: .bold, design: .rounded))
                .kerning(-0.2)
                .foregroundColor(.vText)
                .privacyMasked(maskLevel.hidesAmount)

            HStack(spacing: 3) {
                Image(systemName: pair.isUp ? "arrowtriangle.up.fill" : "arrowtriangle.down.fill")
                    .font(.system(size: 8))
                Text("\(absDelta)%")
                    .font(.system(size: 11, weight: .medium))
            }
            .foregroundColor(pair.isUp ? .vGreen : .vRed)
            .privacyMasked(maskLevel.hidesPercent)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.horizontal, 14)
        .padding(.vertical, 12)
        .background(Color.white.opacity(0.04))
        .clipShape(RoundedRectangle(cornerRadius: 14))
        .overlay(
            RoundedRectangle(cornerRadius: 14)
                .stroke(Color.white.opacity(0.07), lineWidth: 1)
        )
    }

    private var absDelta: String {
        let f = NumberFormatter()
        f.locale = Locale(identifier: "de_DE")
        f.minimumFractionDigits = 2
        f.maximumFractionDigits = 2
        return f.string(from: NSNumber(value: Swift.abs(pair.deltaPct))) ?? ""
    }

    private func formatRate(_ rate: Double) -> String {
        let f = NumberFormatter()
        f.locale = Locale(identifier: "de_DE")
        if rate >= 1000      { f.maximumFractionDigits = 0 }
        else if rate >= 100  { f.maximumFractionDigits = 2 }
        else if rate >= 1    { f.maximumFractionDigits = 4 }
        else                 { f.maximumFractionDigits = 6 }
        return f.string(from: NSNumber(value: rate)) ?? "\(rate)"
    }
}

// MARK: - Preview ────────────────────────────────────────────────────────────

struct HomeView_Previews: PreviewProvider {
    static var previews: some View {
        HomeView()
            .preferredColorScheme(.dark)
    }
}
