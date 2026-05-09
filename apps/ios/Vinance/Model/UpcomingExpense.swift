//
//  UpcomingExpense.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import Foundation

struct UpcomingExpense: Identifiable {
    let id = UUID()
    let name: String
    let amount: Double
    let daysRemaining: Int
}
