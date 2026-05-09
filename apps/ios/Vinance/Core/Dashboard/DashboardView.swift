//
//  DashboardView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

struct DashboardView: View {
    @State private var selectedView: PortfolioView = .savings
    @State private var isBalanceHidden = false
    @State private var showProfileView = false
    @State private var selectedTimePeriod: TimePeriod = .week1
    @EnvironmentObject var authViewModel: AuthViewModel
    
    // Mock data for dashboard
    let availableBudget = 650
    let totalBudget = 1580
    let underBudget = 12
    
    // Chart data points
    let spendingChartData: [ChartPoint] = [
        ChartPoint(value: 500, date: Date().addingTimeInterval(-30 * 24 * 60 * 60)),
        ChartPoint(value: 650, date: Date().addingTimeInterval(-25 * 24 * 60 * 60)),
        ChartPoint(value: 600, date: Date().addingTimeInterval(-20 * 24 * 60 * 60)),
        ChartPoint(value: 700, date: Date().addingTimeInterval(-15 * 24 * 60 * 60)),
        ChartPoint(value: 750, date: Date().addingTimeInterval(-10 * 24 * 60 * 60)),
        ChartPoint(value: 900, date: Date().addingTimeInterval(-5 * 24 * 60 * 60)),
        ChartPoint(value: 850, date: Date())
    ]
    
    let netWorthChartData: [ChartPoint] = [
        ChartPoint(value: 10000, date: Date().addingTimeInterval(-30 * 24 * 60 * 60)),
        ChartPoint(value: 10200, date: Date().addingTimeInterval(-25 * 24 * 60 * 60)),
        ChartPoint(value: 10150, date: Date().addingTimeInterval(-20 * 24 * 60 * 60)),
        ChartPoint(value: 10300, date: Date().addingTimeInterval(-15 * 24 * 60 * 60)),
        ChartPoint(value: 10500, date: Date().addingTimeInterval(-10 * 24 * 60 * 60)),
        ChartPoint(value: 10800, date: Date().addingTimeInterval(-5 * 24 * 60 * 60)),
        ChartPoint(value: 11048, date: Date())
    ]
    
    // Mock transactions
    let recentTransactions = [
        Transaction(name: "Peloton", amount: 12.99, category: .fitness),
        Transaction(name: "Brooklyn Animal Sh...", amount: 50.00, category: .donations),
        Transaction(name: "Starbucks", amount: 7.08, category: .coffee)
    ]
    
    // Mock budget categories
    let budgetCategories = [
        BudgetCategory(name: "Travel", icon: "airplane", color: .orange, spent: 24.80, remaining: -24.80),
        BudgetCategory(name: "Shopping", icon: "bag", color: .pink, spent: 0, remaining: 62.23),
        BudgetCategory(name: "Food", icon: "fork.knife", color: .yellow, spent: 0, remaining: 48.00),
        BudgetCategory(name: "Health", icon: "heart", color: .green, spent: 0, remaining: 154.00),
        BudgetCategory(name: "Entertainment", icon: "tv", color: .blue, spent: 0, remaining: 30.00)
    ]
    
    // Mock upcoming expenses
    let upcomingExpenses = [
        UpcomingExpense(name: "Insurance", amount: 1000.00, daysRemaining: 7),
        UpcomingExpense(name: "Gym", amount: 550.00, daysRemaining: 9),
//        UpcomingExpense(name: "AC Installment", amount: 354.00, daysRemaining: 11)
    ]
    
    enum PortfolioView {
        case savings
        case portfolio
    }
    
    var body: some View {
        let currentUser = authViewModel.currentUser
        
        VStack(spacing: 0) {
            // Header section
            MainHeader(currentUser: currentUser, isBalanceHidden: isBalanceHidden)
            
            // Main section
            ScrollView {
                VStack(spacing: 20) {
                    // Charts section with paging scroll effect
                    GeometryReader { geometry in
                        TabView {
                            NetworthChart(
                                netWorthChartData: netWorthChartData,
                                netWorthValue: "$11,048",
                                netWorthGrowth: "1.47%",
                                isNetWorthPositive: true
                            )
                            .frame(width: geometry.size.width - 32)
                            .padding(.horizontal, 16)
                            
                            SpendingChart(
                                selectedTimePeriod: selectedTimePeriod,
                                spendingChartData: spendingChartData,
                                spendingValue: "$566",
                                underBudgetValue: 74,
                                colors: [.orange, .green]
                            )
                            .frame(width: geometry.size.width - 32)
                            .padding(.horizontal, 16)
                        }
                        .tabViewStyle(PageTabViewStyle(indexDisplayMode: .never))
                    }
                    .frame(height: 220) // Set a fixed height for the charts
                    
                    // Budget summary card - original implementation
                    BudgetSummaryView(
                        availableBudget: availableBudget,
                        totalBudget: totalBudget,
                        underBudget: underBudget
                    )
                    .padding(.horizontal)
                    
                    // Transactions section
                    VStack(alignment: .leading, spacing: 12) {
                        HStack {
                            Text("TO REVIEW")
                                .font(.caption)
                                .fontWeight(.medium)
                                .foregroundColor(.secondary)
                            
                            Spacer()
                            
                            Button(action: {
                                // View all transactions
                            }) {
                                HStack(spacing: 4) {
                                    Text("view all")
                                        .font(.caption)
                                    Image(systemName: "chevron.right")
                                        .font(.system(size: 8))
                                }
                                .foregroundColor(.blue)
                            }
                        }
                        
                        // Transaction list
                        ForEach(recentTransactions) { transaction in
                            TransactionItemView(transaction: transaction)
                                .padding(.vertical, 4)
                        }
                        
                        Button(action: {
                            // Mark all transactions as reviewed
                        }) {
                            Text("MARK AS REVIEWED")
                                .font(.caption)
                                .fontWeight(.medium)
                                .foregroundColor(.blue)
                                .frame(maxWidth: .infinity)
                                .padding(.vertical, 10)
                        }
                        .background(Color.blue.opacity(0.1))
                        .cornerRadius(8)
                    }
                    .padding(.horizontal)
                    
//                    // Budget categories section
//                    VStack(alignment: .leading, spacing: 12) {
//                        HStack {
//                            Text("BUDGETS")
//                                .font(.caption)
//                                .fontWeight(.medium)
//                                .foregroundColor(.secondary)
//
//                            Spacer()
//
//                            Button(action: {
//                                // View all budgets
//                            }) {
//                                HStack(spacing: 4) {
//                                    Text("view all")
//                                        .font(.caption)
//                                    Image(systemName: "chevron.right")
//                                        .font(.system(size: 8))
//                                }
//                                .foregroundColor(.blue)
//                            }
//                        }
//
//                        // Budget categories grid
//                        HStack(spacing: 12) {
//                            ForEach(budgetCategories) { category in
//                                BudgetCategoryView(category: category)
//                            }
//                        }
//                        .padding(.vertical, 4)
//                    }
//                    .padding(.horizontal)
                    
                    // Upcoming expenses section
                    VStack(alignment: .leading, spacing: 12) {
                        HStack {
                            Text("UPCOMING")
                                .font(.caption)
                                .fontWeight(.medium)
                                .foregroundColor(.secondary)
                            
                            Spacer()
                            
                            Button(action: {
                                // View all upcoming expenses
                            }) {
                                HStack(spacing: 4) {
                                    Text("view all")
                                        .font(.caption)
                                    Image(systemName: "chevron.right")
                                        .font(.system(size: 8))
                                }
                                .foregroundColor(.blue)
                            }
                        }
                        
                        // Upcoming expenses list
                        HStack(spacing: 20) {
                            ForEach(upcomingExpenses) { expense in
                                UpcomingExpenseView(expense: expense)
                            }
                        }
                        .padding(.vertical, 4)
                    }
                    .padding(.horizontal)
                    
                    // Income section header
                    HStack {
                        Text("INCOME")
                            .font(.caption)
                            .fontWeight(.medium)
                            .foregroundColor(.secondary)
                        
                        Spacer()
                        
                        Button(action: {
                            // View all income
                        }) {
                            HStack(spacing: 4) {
                                Text("view all")
                                    .font(.caption)
                                Image(systemName: "chevron.right")
                                    .font(.system(size: 8))
                            }
                            .foregroundColor(.blue)
                        }
                    }
                    .padding(.horizontal)
                }
                .padding(.vertical)
            }
            .background(Color.lightBackground.ignoresSafeArea())
        }
        .navigationBarHidden(true)
    }
}

struct DashboardView_Previews: PreviewProvider {
    static var previews: some View {
        DashboardView()
    }
}
