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
    @EnvironmentObject var authViewModel: AuthViewModel
    
    // Mock data - will be replaced with API calls later
    let username = "Vincent"
    
    let savingsCurrency = "$"
    let savingsValue = "12.456,67"
    
    let portfolioCurrency = "€"
    let portfolioValue = "25.000,00"
    
    enum PortfolioView {
        case savings
        case portfolio
    }
    
    func getBalance() -> String {
        let selectedCurrency = selectedView == .savings ? savingsCurrency : portfolioCurrency
        let selectedValue = selectedView == .savings ? savingsValue : portfolioValue
        
        return "\(selectedCurrency) \(selectedValue)"
    }
    
    var body: some View {
        let currentUser = authViewModel.currentUser
        
        VStack(spacing: 0) {
            // Header section
            HStack(alignment: .top) {
                // Left side of header
                VStack(alignment: .leading, spacing: 4) {
                    Text("Hi, \(currentUser?.username ?? "User")")
                        .font(.headline)
                        .foregroundColor(Color(hex: 0x393E46))
                    
                    // Toggle buttons
                    HStack(spacing: 12) {
                        Button(action: {
                            selectedView = .savings
                        }) {
                            Text("Savings")
                                .font(.system(size: 16, weight: .semibold))
                                .foregroundColor(selectedView == .savings ? Color(hex: 0x393E46) : Color(hex: 0x393E46).opacity(0.5))
                        }
                        
                        Button(action: {
                            selectedView = .portfolio
                        }) {
                            Text("Portfolio")
                                .font(.system(size: 16, weight: .semibold))
                                .foregroundColor(selectedView == .portfolio ? Color(hex: 0x393E46) : Color(hex: 0x393E46).opacity(0.5))
                        }
                    }
                    
                    // Account value - changes based on selected view
                    HStack(alignment: .center, spacing: 8) {
                        Text("\(isBalanceHidden ? "***" : getBalance())")
                            .font(.system(size: 24, weight: .bold))
                            .foregroundColor(Color(hex: 0x393E46))
                        
                        Button(action: {
                            isBalanceHidden.toggle()
                        }) {
                            Image(systemName: isBalanceHidden ? "eye.slash" : "eye")
                                .font(.system(size: 16))
                                .foregroundColor(Color(hex: 0x393E46).opacity(0.7))
                        }
                    }
                    .padding(.top, 4)
                }
                
                Spacer()
                
                // Right side of header - profile icon
                Button(action: {
                    showProfileView = true
                }) {
                    Image(systemName: "person")
                        .font(.system(size: 18))
                }
                .sheet(isPresented: $showProfileView) {
                    ProfileView()
                }
            }
            .padding([.horizontal, .top])
            
            // Content based on selected view
            ScrollView {
                VStack(alignment: .leading, spacing: 16) {
                    Text("Vinance Dashboard")
                        .font(.system(size: 24))
                        .fontWeight(.bold)
                        .foregroundColor(Color(hex: 0x393E46))
                    
                    // More content can be added here
                    Spacer()
                }
                .frame(maxWidth: .infinity)
            }
        }
        .navigationBarHidden(true)
    }
}

struct DashboardView_Previews: PreviewProvider {
    static var previews: some View {
        DashboardView()
    }
}
