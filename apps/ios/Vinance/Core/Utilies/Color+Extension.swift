//
//  Color+Extension.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import Foundation
import SwiftUI

extension Color {

    // ── Legacy colors (kept for compatibility) ──────────────────────────────
    static let darkText        = Color(red: 57/255, green: 62/255, blue: 70/255)
    static let darkBlue        = Color(red: 10/255, green: 20/255, blue: 40/255)
    static let lightBackground = Color(red: 245/255, green: 245/255, blue: 245/255)

    // ── Vinance Design System — dark theme ──────────────────────────────────
    static let vBg       = Color(hex: 0x0a0a0f)  // App background
    static let vSurface  = Color(hex: 0x13131a)  // Card / bottom sheet
    static let vSurface2 = Color(hex: 0x1c1c26)  // Secondary surface
    static let vSurface3 = Color(hex: 0x242432)  // Tertiary surface / divider bg
    static let vText     = Color(hex: 0xf0f0f8)  // Primary text
    static let vText2    = Color(hex: 0x8888aa)  // Secondary text
    static let vText3    = Color(hex: 0x55556a)  // Muted label text
    static let vAccent   = Color(hex: 0x7c6aff)  // Purple accent
    static let vAccent2  = Color(hex: 0xa594ff)  // Lighter purple accent
    static let vGreen    = Color(hex: 0x34d399)  // Income / positive
    static let vRed      = Color(hex: 0xf87171)  // Expense / negative
    static let vYellow   = Color(hex: 0xfbbf24)  // Warning / highlight
    static let vBlue     = Color(hex: 0x60a5fa)  // Info / savings
    static let vPink     = Color(hex: 0xf472b6)  // Accent pink
    static let vOrange   = Color(hex: 0xfb923c)  // Accent orange

    // ── Hex integer initialiser ─────────────────────────────────────────────
    init(hex: UInt, alpha: Double = 1.0) {
        self.init(
            .sRGB,
            red:     Double((hex >> 16) & 0xFF) / 255.0,
            green:   Double((hex >> 8)  & 0xFF) / 255.0,
            blue:    Double( hex        & 0xFF) / 255.0,
            opacity: alpha
        )
    }
}
