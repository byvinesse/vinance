//
//  HomeViewTests.swift
//  VinanceTests
//
//  Unit tests for the value types & formatting helpers backing HomeView
//  (VNC-38). The view layout itself is exercised via SwiftUI Previews;
//  these tests cover the deterministic, locale-sensitive logic.
//

import XCTest
import SwiftUI
@testable import Vinance

final class HomeViewTests: XCTestCase {

    // MARK: - MaskLevel ─────────────────────────────────────────────────

    func test_maskLevel_cyclesVisibleAmountAllVisible() {
        XCTAssertEqual(MaskLevel.visible.next,      .amountHidden)
        XCTAssertEqual(MaskLevel.amountHidden.next, .allHidden)
        XCTAssertEqual(MaskLevel.allHidden.next,    .visible)
    }

    func test_maskLevel_hidesAmountFromLevel1Onwards() {
        XCTAssertFalse(MaskLevel.visible.hidesAmount)
        XCTAssertTrue (MaskLevel.amountHidden.hidesAmount)
        XCTAssertTrue (MaskLevel.allHidden.hidesAmount)
    }

    func test_maskLevel_hidesPercentOnlyAtAllHidden() {
        XCTAssertFalse(MaskLevel.visible.hidesPercent)
        XCTAssertFalse(MaskLevel.amountHidden.hidesPercent)
        XCTAssertTrue (MaskLevel.allHidden.hidesPercent)
    }

    func test_maskLevel_labelsMatchDesign() {
        XCTAssertEqual(MaskLevel.visible.label,      "Visible")
        XCTAssertEqual(MaskLevel.amountHidden.label, "฿ Hidden")
        XCTAssertEqual(MaskLevel.allHidden.label,    "All Hidden")
    }

    // MARK: - MoneyFormatter (EU style) ─────────────────────────────────

    func test_split_thb_emitsCommaDecimalAndPeriodGrouping() {
        let p = MoneyFormatter.split(284_520, in: .thb)
        XCTAssertEqual(p.symbol,  "฿")
        XCTAssertEqual(p.integer, "284.520")
        XCTAssertEqual(p.decimal, "00")
    }

    func test_split_thb_keepsTwoDecimalDigits() {
        let p = MoneyFormatter.split(840.5, in: .thb)
        XCTAssertEqual(p.integer, "840")
        XCTAssertEqual(p.decimal, "50")
    }

    func test_split_idr_omitsDecimalSegment() {
        let p = MoneyFormatter.split(12_500_000, in: .idr)
        XCTAssertEqual(p.symbol,  "Rp")
        XCTAssertEqual(p.integer, "12.500.000")
        XCTAssertEqual(p.decimal, "")
    }

    func test_split_handlesNegativeValuesAsAbsolute() {
        // The sign is added by `formatted(_:in:sign:)`; `split` always
        // returns the magnitude so callers don't double-up minus signs.
        let p = MoneyFormatter.split(-840, in: .thb)
        XCTAssertEqual(p.integer, "840")
    }

    func test_formatted_appliesUnicodeMinusForExpenses() {
        let s = MoneyFormatter.formatted(840, in: .thb, sign: .minus)
        XCTAssertEqual(s, "−฿840,00")          // U+2212
        XCTAssertTrue(s.contains("\u{2212}"))
        XCTAssertFalse(s.contains("-"))        // never the ASCII hyphen
    }

    func test_formatted_appliesPlusForIncome() {
        let s = MoneyFormatter.formatted(68_000, in: .thb, sign: .plus)
        XCTAssertEqual(s, "+฿68.000,00")
    }

    func test_formatted_emitsBareAmountForTransfers() {
        let s = MoneyFormatter.formatted(5_000, in: .thb, sign: .none)
        XCTAssertEqual(s, "฿5.000,00")
    }

    func test_formatted_idrAccountHasNoDecimalTail() {
        let s = MoneyFormatter.formatted(12_500_000, in: .idr, sign: .none)
        XCTAssertEqual(s, "Rp12.500.000")
    }

    // MARK: - HomeTxItem ────────────────────────────────────────────────

    func test_homeTxItem_signMapsToTransactionKind() {
        let exp = HomeTxItem(kind: .expense,  categoryName: "Food", categoryColor: .vRed,
                             emoji: "🍜", account: "K", toAccount: nil, description: "",
                             amount: 1, currency: .thb, time: "00:00", dateKey: "today")
        let inc = HomeTxItem(kind: .income,   categoryName: "Salary", categoryColor: .vGreen,
                             emoji: "💼", account: "B", toAccount: nil, description: "",
                             amount: 1, currency: .thb, time: "00:00", dateKey: "today")
        let xfr = HomeTxItem(kind: .transfer, categoryName: "Transfer", categoryColor: .vBlue,
                             emoji: "🔄", account: "K", toAccount: "B", description: "",
                             amount: 1, currency: .thb, time: "00:00", dateKey: "today")

        if case .minus = exp.sign {} else { XCTFail("expense should be .minus") }
        if case .plus  = inc.sign {} else { XCTFail("income should be .plus")  }
        if case .none  = xfr.sign {} else { XCTFail("transfer should be .none") }
    }

    func test_homeTxItem_formattedAmount_attachesSign() {
        let groceries = HomeTxItem(kind: .expense, categoryName: "Groceries",
                                   categoryColor: .vYellow, emoji: "🛒",
                                   account: "Kasikorn Bank", toAccount: nil,
                                   description: "Tops Supermarket",
                                   amount: 840, currency: .thb,
                                   time: "11:24", dateKey: "today")
        XCTAssertEqual(groceries.formattedAmount, "−฿840,00")
    }

    // MARK: - DateGroup labels ──────────────────────────────────────────

    func test_dateGroup_labelsTodayAndYesterday() {
        // The grouping is exercised through MockHomeData. We assert two
        // anchors here so a label-key drift (e.g. "today" → "Today!")
        // breaks loudly rather than silently passing through.
        let txs = MockHomeData.transactions
        XCTAssertTrue(txs.contains { $0.dateKey == "today"     })
        XCTAssertTrue(txs.contains { $0.dateKey == "yesterday" })
    }

    // MARK: - Mock data shape ───────────────────────────────────────────

    func test_mockHomeData_includesMultiCurrencyAccount() {
        let codes = Set(MockHomeData.accounts.map { $0.currency.code })
        XCTAssertTrue(codes.contains("THB"))
        XCTAssertTrue(codes.contains("IDR"))   // BCA Mandiri credit card
    }

    func test_mockHomeData_fxGridHasFourPairsAgainstTHB() {
        XCTAssertEqual(MockHomeData.fxPairs.count, 4)
        XCTAssertTrue(MockHomeData.fxPairs.allSatisfy { $0.to == "THB" })
    }

    func test_fxPair_isUpReflectsDeltaSign() {
        XCTAssertTrue (FxPair(from: "USD", to: "THB", rate: 1, deltaPct:  0.10).isUp)
        XCTAssertFalse(FxPair(from: "USD", to: "THB", rate: 1, deltaPct: -0.10).isUp)
        XCTAssertTrue (FxPair(from: "USD", to: "THB", rate: 1, deltaPct:  0.00).isUp,
                       "zero delta is treated as up so the indicator never blanks")
    }
}
