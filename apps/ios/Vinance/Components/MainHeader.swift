//
//  MainHeader.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct MainHeader: View {
    
    var currentUser: User?
    @State var selectedView: PortfolioView = .savings
    @State var isBalanceHidden: Bool
    @State var showProfileView = false
    
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
        HStack(alignment: .top) {
            // Left side of header
            VStack(alignment: .leading, spacing: 4) {
                Text("Hi, \(currentUser?.username ?? "User")")
                    .font(.headline)
                    .foregroundColor(Color.darkText)
                
                // Toggle buttons
                HStack(spacing: 12) {
                    Button(action: {
                        selectedView = .savings
                    }) {
                        Text("Savings")
                            .font(.system(size: 16, weight: .semibold))
                            .foregroundColor(selectedView == .savings ? Color.darkText : Color.darkText.opacity(0.5))
                    }
                    
                    Button(action: {
                        selectedView = .portfolio
                    }) {
                        Text("Portfolio")
                            .font(.system(size: 16, weight: .semibold))
                            .foregroundColor(selectedView == .portfolio ? Color.darkText : Color.darkText.opacity(0.5))
                    }
                }
                
                // Account value - changes based on selected view
                HStack(alignment: .center, spacing: 8) {
                    Text("\(isBalanceHidden ? "***" : getBalance())")
                        .font(.system(size: 24, weight: .bold))
                        .foregroundColor(Color.darkText)
                    
                    Button(action: {
                        isBalanceHidden.toggle()
                    }) {
                        Image(systemName: isBalanceHidden ? "eye.slash" : "eye")
                            .font(.system(size: 16))
                            .foregroundColor(Color.darkText.opacity(0.7))
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
    }
}
