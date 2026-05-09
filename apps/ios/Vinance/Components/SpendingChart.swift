//
//  MonthlySpendingChart.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct SpendingChart: View {
    
    @State var selectedTimePeriod: TimePeriod
    let spendingChartData: [ChartPoint]
    let spendingValue: String
    let underBudgetValue: Int32
    let colors: [Color]
    
    func getSpendingHorizonTitle() -> String {
        if selectedTimePeriod == TimePeriod.day1 {
            return "Daily Spending"
        }
        if selectedTimePeriod == TimePeriod.week1 {
            return "Weekly Spending"
        }
        if selectedTimePeriod == TimePeriod.month1 {
            return "Monthly Spending"
        }
        if selectedTimePeriod == TimePeriod.month3 {
            return "Last 3 Months Spending"
        }
        return "Year to Date Spending"
    }
    
    func isShowUnderBudget() -> Bool {
        if underBudgetValue == 0 {
            return false
        }
        return true
    }
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text(getSpendingHorizonTitle())
                    .font(.headline)
                    .foregroundColor(.white)

                Spacer()

                Button(action: {
                    // View transactions
                }) {
                    HStack(spacing: 4) {
                        Text("TRANSACTIONS")
                            .font(.caption)
                            .fontWeight(.medium)
                        Image(systemName: "chevron.right")
                            .font(.system(size: 8))
                    }
                    .foregroundColor(.gray)
                }
            }

            Text(spendingValue)
                .font(.title2)
                .fontWeight(.bold)
                .foregroundColor(.white)

            // Line chart
            ZStack(alignment: .bottomLeading) {
                // Chart
                LineChartView(
                    data: spendingChartData,
                    showUnderBudget: isShowUnderBudget(),
                    underBudgetValue: "\(underBudgetValue)",
                    colors: colors
                )
                .frame(height: 60)
            }
            .padding(.top, 8)
            
            // Time period selection
            HStack(spacing: 12) {
                ForEach(TimePeriod.allCases) { period in
                    Button(action: {
                        selectedTimePeriod = period
                    }) {
                        Text(period.rawValue)
                            .font(.caption2)
                            .padding(.vertical, 4)
                            .padding(.horizontal, 8)
                            .background(
                                selectedTimePeriod == period ?
                                Color.blue.opacity(0.8) : Color.gray.opacity(0.3)
                            )
                            .foregroundColor(
                                selectedTimePeriod == period ? .white : .gray
                            )
                            .cornerRadius(12)
                    }
                }
            }
            .padding(.top, 8)
        }
        .padding()
        .background(Color.darkBlue)
        .cornerRadius(12)
        .frame(maxWidth: .infinity)
    }
}
