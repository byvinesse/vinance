//
//  BudgetSummaryView.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct BudgetSummaryView: View {
    let availableBudget: Int
    let totalBudget: Int
    let underBudget: Int
    
    func getBudgetTitle() -> String {
        if availableBudget < totalBudget {
            return "$\(availableBudget) left"
        }
        return "$\(availableBudget) over"
    }
    
    func getRemainingBudgetTitle() -> String {
        if availableBudget < totalBudget {
            return "$\(underBudget) under"
        }
        return "$\(underBudget) over"
    }
    
    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text(getBudgetTitle())
                    .font(.title2)
                    .fontWeight(.bold)
                    .foregroundColor(.white)
                
                Spacer()
                
                Button(action: {
                    // Show budget info
                }) {
                    Image(systemName: "info.circle")
                        .foregroundColor(.white.opacity(0.7))
                }
            }
            
            Text("out of $\(totalBudget) budgeted")
                .font(.caption)
                .foregroundColor(.white.opacity(0.8))
            
            // Budget progress bar
            ZStack(alignment: .leading) {
                RoundedRectangle(cornerRadius: 2)
                    .frame(height: 4)
                    .foregroundColor(.white.opacity(0.3))
                
                // Progress bar (70% of width for example)
                RoundedRectangle(cornerRadius: 2)
                    .frame(width: UIScreen.main.bounds.width * 0.7 - 40, height: 4)
                    .foregroundColor(.green)
            }
            .padding(.top, 4)
            
            // Under budget indicator
            HStack {
                Spacer()
                
                Text(getRemainingBudgetTitle())
                    .font(.caption)
                    .fontWeight(.bold)
                    .padding(6)
                    .background(availableBudget < totalBudget ? Color.green : Color.red)
                    .foregroundColor(.white)
                    .clipShape(RoundedRectangle(cornerRadius: 4))
            }
            .padding(.top, -10)
        }
        .padding()
        .background(
            LinearGradient(
                gradient: Gradient(colors: [Color.blue.opacity(0.8), Color.blue.opacity(0.6)]),
                startPoint: .topLeading,
                endPoint: .bottomTrailing
            )
        )
        .clipShape(RoundedRectangle(cornerRadius: 12))
    }
}

