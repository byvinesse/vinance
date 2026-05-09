//
//  InputView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

struct InputView: View {
    @Binding var text: String
    let title: String
    let placeholder: String
    var isSecureField = false
    
    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(title)
                .foregroundColor(Color(.darkGray))
                .fontWeight(.semibold)
                .font(.system(size: 20))
            
            if isSecureField {
                SecureField(placeholder, text: $text)
                    .font(.system(size: 20))
            } else {
                TextField(placeholder, text: $text)
                    .font(.system(size: 20))
            }
            
            Divider()
        }
    }
}

struct InputView_Previews: PreviewProvider {
    static var previews: some View {
        InputView(text: .constant(""), title: "Email", placeholder: "name@example.com")
    }
}
