//
//  BudgetCategoryVie.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct BudgetCategoryView: View {
    let category: BudgetCategory
    
    var body: some View {
        VStack(spacing: 8) {
            ZStack {
                Circle()
                    .fill(category.color.opacity(0.2))
                    .frame(width: 50, height: 50)
                
                Image(systemName: category.icon)
                    .font(.system(size: 20))
                    .foregroundColor(category.color)
            }
            
            Text("$\(String(format: "%.2f", abs(category.remaining)))")
                .font(.caption)
                .fontWeight(.bold)
            
            Text(category.remaining < 0 ? "over" : "left")
                .font(.system(size: 10))
                .foregroundColor(.secondary)
        }
        .frame(maxWidth: .infinity)
    }
}
