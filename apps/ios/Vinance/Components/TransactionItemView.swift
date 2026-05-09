//
//  TransactionItemRow.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct TransactionItemView: View {
    let transaction: Transaction
    
    var body: some View {
        HStack(spacing: 12) {
            Image(systemName: "checkmark.square")
                .foregroundColor(.gray)
            
            Text(transaction.name)
                .font(.subheadline)
            
            Spacer()
            
            // Category Tag
            Text(transaction.category.rawValue.uppercased())
                .font(.system(size: 10))
                .fontWeight(.bold)
                .padding(.horizontal, 8)
                .padding(.vertical, 4)
                .background(transaction.category.color.opacity(0.2))
                .foregroundColor(transaction.category.color)
                .clipShape(RoundedRectangle(cornerRadius: 4))
            
            Text("$\(String(format: "%.2f", transaction.amount))")
                .font(.subheadline)
                .fontWeight(.medium)
        }
    }
}
