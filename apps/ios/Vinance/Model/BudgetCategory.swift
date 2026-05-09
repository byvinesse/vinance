//
//  BudgetCategory.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import Foundation
import SwiftUI

struct BudgetCategory: Identifiable {
    let id = UUID()
    let name: String
    let icon: String
    let color: Color
    let spent: Double
    let remaining: Double
}
