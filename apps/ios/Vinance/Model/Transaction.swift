//
//  Transaction.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import Foundation
import SwiftUI

struct Transaction: Identifiable {
    let id = UUID()
    let name: String
    let amount: Double
    let category: Category
    
    enum Category: String {
        case fitness = "fitness"
        case donations = "donations"
        case coffee = "coffee"
        
        var color: Color {
            switch self {
            case .fitness:
                return .yellow
            case .donations:
                return .orange
            case .coffee:
                return .brown
            }
        }
    }
}
