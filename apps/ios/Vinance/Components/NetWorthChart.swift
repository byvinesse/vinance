//
//  NetWorthChart.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct NetworthChart: View {
    
    let netWorthChartData: [ChartPoint]
    let netWorthValue: String
    let netWorthGrowth: String
    let isNetWorthPositive: Bool
    
    func getNetWorthGrowthArrow() -> String {
        if isNetWorthPositive {
            return "arrow.up"
        }
        return "arrow.down"
    }
    
    func getChartColors() -> [Color] {
        if isNetWorthPositive {
            return [.yellow, .green]
        }
        return [.orange, .red]
    }
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text("Net Worth")
                    .font(.headline)
                    .foregroundColor(.white)
                
                Spacer()
                
                Button(action: {
                    // View transactions
                }) {
                    HStack(spacing: 4) {
                        Text("Accounts")
                            .font(.caption)
                            .fontWeight(.medium)
                        Image(systemName: "chevron.right")
                            .font(.system(size: 8))
                    }
                    .foregroundColor(.gray)
                }
            }
            
            Text(netWorthValue)
                .font(.title2)
                .fontWeight(.bold)
                .foregroundColor(.white)
            
            HStack {
                Image(systemName: getNetWorthGrowthArrow())
                    .foregroundColor(.green)
                Text(netWorthGrowth)
                    .font(.caption)
                    .foregroundColor(.green)
            }
            
            // Line chart for net worth
            ZStack(alignment: .bottomLeading) {
                LineChartView(
                    data: netWorthChartData,
                    showUnderBudget: false,
                    colors: getChartColors()
                )
                .frame(height: 60)
            }
            .padding(.top, 8)
        }
        .padding()
        .background(Color.darkBlue)
        .cornerRadius(12)
        .frame(maxWidth: .infinity)
    }
}
