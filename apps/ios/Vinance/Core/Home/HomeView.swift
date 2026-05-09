//
//  HomeView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/05/26.
//
//  Reference design: finance_app_2.html (VNC-36)
//  Dark-theme personal finance home screen with:
//    • Greeting header + balance-visibility toggle
//    • Total balance card (฿) with income / expense / saved pills
//    • Quick-action buttons (Expense, Income, Transfer, Report)
//    • Horizontal-scrolling account cards
//    • Grouped recent-transaction list
//

import SwiftUI

// MARK: - Local Model Types ─────────────────────────────────────────────────

/// Lightweight account model used only on the home screen (first iteration).
struct HomeAccount: Identifiable {
    let id = UUID()
    let bank: String
    let amount: Double
    let accountName: String
    let gradient: [Color]
    let dotColor: Color

    var formattedAmount: String { Self.baht(amount) }

    static func baht(_ v: Double) -> String {
        let f = NumberFormatter()
        f.numberStyle = .decimal
        f.groupingSeparator = ","
        f.maximumFractionDigits = 0
        return "฿" + (f.string(from: NSNumber(value: v)) ?? "\(Int(v))")
    }
}

/// Transaction display type for the home screen.
enum HomeTxType { case expense, income, transfer }

/// Lightweight transaction model used only on the home screen (first iteration).
struct HomeTxItem: Identifiable {
    let id = UUID()
    let category: String
    let account: String
    let description: String
    let amount: Double
    let emoji: String
    let time: String
    let rawDate: String          // e.g. "02 May"
    let type: HomeTxType
    let tint: Color

    var dateLabel: String {
        switch rawDate {
        case "02 May": return "Today"
        case "01 May": return "Yesterday"
        default:       return rawDate
        }
    }

    var formattedAmount: String {
        let f = NumberFormatter()
        f.numberStyle = .decimal
        f.groupingSeparator = ","
        f.maximumFractionDigits = 0
        let n = f.string(from: NSNumber(value: amount)) ?? "\(Int(amount))"
        switch type {
        case .expense:  return "−฿\(n)"
        case .income:   return "+฿\(n)"
        case .transfer: return "฿\(n)"
        }
    }

    var amountColor: Color {
        switch type {
        case .expense:  return .vRed
        case .income:   return .vGreen
        case .transfer: return .vText
        }
    }
}

// MARK: - Mock Data ──────────────────────────────────────────────────────────

private enum MockHomeData {

    static let accounts: [HomeAccount] = [
        HomeAccount(
            bank: "Kasikorn Bank",
            amount: 142_300,
            accountName: "Savings Account",
            gradient: [Color(hex: 0x1e1b4b), Color(hex: 0x312e81)],
            dotColor: Color(hex: 0xf59e0b)
        ),
        HomeAccount(
            bank: "Bangkok Bank",
            amount: 89_220,
            accountName: "Current Account",
            gradient: [Color(hex: 0x064e3b), Color(hex: 0x065f46)],
            dotColor: Color(hex: 0x34d399)
        ),
        HomeAccount(
            bank: "SCB",
            amount: 53_000,
            accountName: "Investment Fund",
            gradient: [Color(hex: 0x1e3a5f), Color(hex: 0x1e3a8a)],
            dotColor: Color(hex: 0x60a5fa)
        )
    ]

    static let transactions: [HomeTxItem] = [
        HomeTxItem(category: "Groceries",     account: "Kasikorn Bank", description: "Tops Supermarket",       amount: 840,    emoji: "🛒", time: "11:24", rawDate: "02 May", type: .expense,  tint: Color(hex: 0xfbbf24, alpha: 0.15)),
        HomeTxItem(category: "Food & Drink",  account: "Kasikorn Bank", description: "Starbucks Siam Paragon", amount: 180,    emoji: "☕", time: "08:50", rawDate: "02 May", type: .expense,  tint: Color(hex: 0x60a5fa, alpha: 0.15)),
        HomeTxItem(category: "Income",        account: "Bangkok Bank",  description: "Monthly Salary",         amount: 68_000, emoji: "💼", time: "09:00", rawDate: "01 May", type: .income,   tint: Color(hex: 0x34d399, alpha: 0.15)),
        HomeTxItem(category: "Subscriptions", account: "SCB",           description: "Spotify Premium",        amount: 99,     emoji: "🎵", time: "00:00", rawDate: "01 May", type: .expense,  tint: Color(hex: 0xf472b6, alpha: 0.15)),
        HomeTxItem(category: "Food & Drink",  account: "Cash",          description: "Or Tor Kor Market",      amount: 320,    emoji: "🍜", time: "13:10", rawDate: "01 May", type: .expense,  tint: Color(hex: 0xfbbf24, alpha: 0.15))
    ]
}

// MARK: - HomeView ───────────────────────────────────────────────────────────

struct HomeView: View {

    @State private var isBalanceHidden = false

    // Summary figures (mock)
    private let totalBalance: Double = 284_520
    private let income:       Double = 68_000
    private let expenses:     Double = 34_280
    private let saved:        Double = 33_720
    private let changeText           = "+4.2%"

    var body: some View {
        ZStack {
            Color.vBg.ignoresSafeArea()

            VStack(spacing: 0) {
                headerBar
                ScrollView(showsIndicators: false) {
                    VStack(spacing: 0) {
                        balanceCard
                            .padding(.top, 18)
                            .padding(.horizontal, 24)

                        quickActionsSection
                            .padding(.top, 4)

                        accountsSection
                            .padding(.top, 4)

                        recentSection

                        // Bottom breathing room above tab bar
                        Color.clear.frame(height: 24)
                    }
                }
            }
        }
        .navigationBarHidden(true)
    }

    // MARK: Header ──────────────────────────────────────────────────────────

    private var headerBar: some View {
        HStack(alignment: .center) {
            VStack(alignment: .leading, spacing: 2) {
                Text(greeting)
                    .font(.system(size: 22, weight: .bold, design: .rounded))
                    .foregroundColor(.vText)
                Text(dateSubtitle)
                    .font(.system(size: 12))
                    .foregroundColor(.vText3)
            }

            Spacer()

            HStack(spacing: 10) {
                // Balance visibility toggle
                Button(action: { withAnimation(.easeInOut(duration: 0.2)) { isBalanceHidden.toggle() } }) {
                    ZStack {
                        Circle()
                            .fill(Color.white.opacity(0.08))
                            .overlay(Circle().stroke(Color.white.opacity(0.14), lineWidth: 1))
                        Image(systemName: isBalanceHidden ? "eye.slash" : "eye")
                            .font(.system(size: 13))
                            .foregroundColor(.vText2)
                    }
                    .frame(width: 34, height: 34)
                }

                // Avatar
                ZStack {
                    Circle()
                        .fill(
                            LinearGradient(
                                colors: [.vAccent, .vPink],
                                startPoint: .topLeading,
                                endPoint: .bottomTrailing
                            )
                        )
                    Text("A")
                        .font(.system(size: 14, weight: .bold))
                        .foregroundColor(.white)
                }
                .frame(width: 36, height: 36)
            }
        }
        .padding(.horizontal, 24)
        .padding(.top, 14)
        .padding(.bottom, 6)
    }

    // MARK: Balance Card ────────────────────────────────────────────────────

    private var balanceCard: some View {
        VStack(alignment: .leading, spacing: 0) {

            Text("TOTAL BALANCE")
                .font(.system(size: 11, weight: .semibold))
                .foregroundColor(.vText3)
                .kerning(1.0)

            // Amount row
            HStack(alignment: .bottom, spacing: 6) {
                Text("฿")
                    .font(.system(size: 22))
                    .foregroundColor(.vText2)
                    .padding(.bottom, 5)

                Text(isBalanceHidden ? "••••••" : formatAmount(totalBalance))
                    .font(.system(size: 40, weight: .bold, design: .rounded))
                    .foregroundColor(.vText)
                    .lineLimit(1)
                    .minimumScaleFactor(0.7)
            }
            .padding(.top, 6)

            // Change indicator
            HStack(spacing: 4) {
                Image(systemName: "arrowtriangle.up.fill")
                    .font(.system(size: 9))
                Text("\(changeText) from last month")
                    .font(.system(size: 12))
            }
            .foregroundColor(.vGreen)
            .padding(.top, 4)

            // Income / Expenses / Saved pills
            HStack(spacing: 10) {
                balancePill(label: "Income",   value: income,   prefix: "+฿", color: .vGreen)
                balancePill(label: "Expenses", value: expenses, prefix: "−฿", color: .vRed)
                balancePill(label: "Saved",    value: saved,    prefix: "฿",  color: .vBlue)
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
                // Purple glow — top right
                RadialGradient(
                    colors: [Color.vAccent.opacity(0.18), Color.clear],
                    center: .topTrailing, startRadius: 0, endRadius: 160
                )
                // Pink glow — bottom left
                RadialGradient(
                    colors: [Color.vPink.opacity(0.08), Color.clear],
                    center: .bottomLeading, startRadius: 0, endRadius: 100
                )
            }
        )
        .clipShape(RoundedRectangle(cornerRadius: 24))
        .overlay(
            RoundedRectangle(cornerRadius: 24)
                .stroke(Color.vAccent.opacity(0.18), lineWidth: 1)
        )
    }

    private func balancePill(label: String, value: Double, prefix: String, color: Color) -> some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(label.uppercased())
                .font(.system(size: 9, weight: .medium))
                .foregroundColor(.vText3)
                .kerning(0.5)

            Text(isBalanceHidden ? "••••" : "\(prefix)\(formatAmount(value))")
                .font(.system(size: 13, weight: .semibold, design: .rounded))
                .foregroundColor(color)
                .lineLimit(1)
                .minimumScaleFactor(0.65)
        }
        .padding(12)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(Color.white.opacity(0.04))
        .clipShape(RoundedRectangle(cornerRadius: 14))
        .overlay(
            RoundedRectangle(cornerRadius: 14)
                .stroke(Color.white.opacity(0.07), lineWidth: 1)
        )
    }

    // MARK: Quick Actions ───────────────────────────────────────────────────

    private var quickActionsSection: some View {
        VStack(alignment: .leading, spacing: 0) {
            sectionHeader("Quick Actions", link: nil)

            HStack(spacing: 10) {
                quickActionBtn(emoji: "💸", label: "Expense",  tint: Color.vRed.opacity(0.16))
                quickActionBtn(emoji: "💰", label: "Income",   tint: Color.vGreen.opacity(0.16))
                quickActionBtn(emoji: "🔄", label: "Transfer", tint: Color.vBlue.opacity(0.16))
                quickActionBtn(emoji: "📊", label: "Report",   tint: Color.vYellow.opacity(0.16))
            }
            .padding(.horizontal, 24)
            .padding(.top, 10)
        }
    }

    private func quickActionBtn(emoji: String, label: String, tint: Color) -> some View {
        VStack(spacing: 8) {
            ZStack {
                RoundedRectangle(cornerRadius: 13)
                    .fill(Color.white.opacity(0.08))
                    .overlay(RoundedRectangle(cornerRadius: 13).fill(tint))
                    .overlay(RoundedRectangle(cornerRadius: 13).stroke(Color.white.opacity(0.14), lineWidth: 1))
                Text(emoji).font(.system(size: 20))
            }
            .frame(width: 42, height: 42)

            Text(label)
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

    // MARK: Accounts ────────────────────────────────────────────────────────

    private var accountsSection: some View {
        VStack(alignment: .leading, spacing: 0) {
            sectionHeader("Accounts", link: "Manage")

            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: 14) {
                    ForEach(MockHomeData.accounts) { account in
                        HomeAccountCard(account: account, isHidden: isBalanceHidden)
                    }
                }
                .padding(.horizontal, 24)
                .padding(.top, 10)
                .padding(.bottom, 4)
            }
        }
    }

    // MARK: Recent Transactions ─────────────────────────────────────────────

    private var recentSection: some View {
        VStack(alignment: .leading, spacing: 0) {
            sectionHeader("Recent", link: "See all")

            let grouped = groupedByDate(MockHomeData.transactions)
            ForEach(grouped, id: \.0) { dateLabel, txs in
                // Date group header
                Text(dateLabel.uppercased())
                    .font(.system(size: 10, weight: .semibold))
                    .foregroundColor(.vText3)
                    .kerning(0.9)
                    .padding(.horizontal, 24)
                    .padding(.top, 14)
                    .padding(.bottom, 4)

                // Transactions in this group
                VStack(spacing: 0) {
                    ForEach(txs) { tx in
                        HomeTxRow(tx: tx, isHidden: isBalanceHidden)
                            .padding(.horizontal, 24)

                        if tx.id != txs.last?.id {
                            Divider()
                                .background(Color.white.opacity(0.07))
                                .padding(.horizontal, 24)
                        }
                    }
                }
            }
        }
    }

    // MARK: Shared Helpers ──────────────────────────────────────────────────

    private func sectionHeader(_ title: String, link: String?) -> some View {
        HStack {
            Text(title)
                .font(.system(size: 16, weight: .semibold, design: .rounded))
                .foregroundColor(.vText)
            Spacer()
            if let link {
                Text(link)
                    .font(.system(size: 13))
                    .foregroundColor(.vAccent2)
            }
        }
        .padding(.horizontal, 24)
        .padding(.top, 22)
    }

    private var greeting: String {
        let h = Calendar.current.component(.hour, from: Date())
        switch h {
        case 0..<12: return "Good morning, Alex 👋"
        case 12..<17: return "Good afternoon, Alex 👋"
        default:     return "Good evening, Alex 👋"
        }
    }

    private var dateSubtitle: String {
        let f = DateFormatter()
        f.dateFormat = "EEEE, d MMM"
        return "\(f.string(from: Date())) · Bangkok"
    }

    private func formatAmount(_ v: Double) -> String {
        let f = NumberFormatter()
        f.numberStyle = .decimal
        f.groupingSeparator = ","
        f.maximumFractionDigits = 0
        return f.string(from: NSNumber(value: v)) ?? "\(Int(v))"
    }

    /// Groups transactions preserving insertion order of date labels.
    private func groupedByDate(_ items: [HomeTxItem]) -> [(String, [HomeTxItem])] {
        var result: [(String, [HomeTxItem])] = []
        var index: [String: Int] = [:]
        for tx in items {
            let lbl = tx.dateLabel
            if let i = index[lbl] {
                result[i].1.append(tx)
            } else {
                index[lbl] = result.count
                result.append((lbl, [tx]))
            }
        }
        return result
    }
}

// MARK: - HomeAccountCard ────────────────────────────────────────────────────

struct HomeAccountCard: View {
    let account: HomeAccount
    let isHidden: Bool

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            // Mastercard-style dots — top right
            HStack {
                Spacer()
                HStack(spacing: 4) {
                    Circle()
                        .fill(account.dotColor)
                        .frame(width: 18, height: 18)
                    Circle()
                        .fill(account.dotColor.opacity(0.4))
                        .frame(width: 18, height: 18)
                }
            }
            .padding(.bottom, 14)

            Text(account.bank)
                .font(.system(size: 10, weight: .semibold))
                .foregroundColor(.white.opacity(0.65))
                .kerning(0.7)

            Text(isHidden ? "฿ ••••••" : account.formattedAmount)
                .font(.system(size: 21, weight: .bold, design: .rounded))
                .foregroundColor(.white)
                .padding(.top, 8)
                .padding(.bottom, 2)
                .lineLimit(1)
                .minimumScaleFactor(0.8)

            Text(account.accountName)
                .font(.system(size: 12))
                .foregroundColor(.white.opacity(0.6))
        }
        .padding(18)
        .frame(width: 196)
        .background(
            LinearGradient(
                colors: account.gradient,
                startPoint: .topLeading,
                endPoint: .bottomTrailing
            )
        )
        .clipShape(RoundedRectangle(cornerRadius: 20))
        .overlay(
            RoundedRectangle(cornerRadius: 20)
                .stroke(Color.white.opacity(0.07), lineWidth: 1)
        )
    }
}

// MARK: - HomeTxRow ──────────────────────────────────────────────────────────

struct HomeTxRow: View {
    let tx: HomeTxItem
    let isHidden: Bool

    var body: some View {
        HStack(spacing: 12) {
            // Glass emoji icon
            ZStack {
                RoundedRectangle(cornerRadius: 14)
                    .fill(Color.white.opacity(0.08))
                    .overlay(RoundedRectangle(cornerRadius: 14).fill(tx.tint))
                    .overlay(RoundedRectangle(cornerRadius: 14).stroke(Color.white.opacity(0.14), lineWidth: 1))
                Text(tx.emoji)
                    .font(.system(size: 20))
            }
            .frame(width: 44, height: 44)

            // Three-row info column
            VStack(alignment: .leading, spacing: 2) {
                Text(tx.category)
                    .font(.system(size: 13, weight: .semibold))
                    .foregroundColor(.vText)
                    .lineLimit(1)

                Text(tx.account)
                    .font(.system(size: 11))
                    .foregroundColor(.vText3)
                    .lineLimit(1)

                Text(tx.description)
                    .font(.system(size: 11).italic())
                    .foregroundColor(.vText2)
                    .lineLimit(1)
            }

            Spacer()

            // Amount + time
            VStack(alignment: .trailing, spacing: 3) {
                Text(isHidden ? "••••" : tx.formattedAmount)
                    .font(.system(size: 14, weight: .semibold, design: .rounded))
                    .foregroundColor(tx.amountColor)

                Text(tx.time)
                    .font(.system(size: 10))
                    .foregroundColor(.vText3)
            }
        }
        .padding(.vertical, 11)
    }
}

// MARK: - Preview ────────────────────────────────────────────────────────────

struct HomeView_Previews: PreviewProvider {
    static var previews: some View {
        HomeView()
            .preferredColorScheme(.dark)
    }
}
