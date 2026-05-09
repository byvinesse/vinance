//
//  UpcomingExpenseView.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct UpcomingExpenseView: View {
    let expense: UpcomingExpense
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text("In \(expense.daysRemaining) days")
                .font(.caption)
                .foregroundColor(.secondary)
            
            HStack(spacing: 8) {
                Image(systemName: expense.name.contains("Insurance") ? "building.columns.fill" : "car.fill")
                    .foregroundColor(.orange)
                    .frame(width: 24, height: 24)
                    .background(Color.orange.opacity(0.2))
                    .clipShape(Circle())
                
                Text(expense.name)
                    .font(.subheadline)
                    .lineLimit(1)
            }
            
            Text("$\(String(format: "%.2f", expense.amount))")
                .font(.subheadline)
                .fontWeight(.semibold)
        }
        .padding(12)
        .frame(width: 150, alignment: .leading)
        .background(Color.white)
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.05), radius: 5, x: 0, y: 2)
    }
}
