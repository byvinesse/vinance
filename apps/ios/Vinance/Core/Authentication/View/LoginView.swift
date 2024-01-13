//
//  LoginView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

extension Color {
    init(hex: UInt, alpha: Double = 1.0) {
        self.init(
            .sRGB,
            red: Double((hex >> 16) & 0xFF) / 255.0,
            green: Double((hex >> 8) & 0xFF) / 255.0,
            blue: Double(hex & 0xFF) / 255.0,
            opacity: alpha
        )
    }
}

struct LoginView: View {
    @State private var email = ""
    @State private var password = ""
    @EnvironmentObject var viewModel: AuthViewModel
    
    var body: some View {
        NavigationStack {
            VStack {
                // app logo
                Image("vinance-logo-bgless")
                    .resizable()
                    .scaledToFill()
                    .frame(width: 344, height: 275)
                
                // hero text
                VStack(alignment: .leading) {
                    Text("Welcome Back")
                        .font(.system(size: 48))
                        .fontWeight(.bold)
                        .foregroundColor(Color(hex: 0x393E46))
                        
                    Text("Sign in to continue")
                        .font(.system(size: 24))
                        .fontWeight(.bold)
                        .foregroundColor(Color(hex: 0x929AAB))
                }
                .padding(.top, -32)
                .padding(.bottom, 48)
                
                // form fields
                VStack(spacing: 24) {
                    InputView(text: $email,
                              title: "Email",
                              placeholder: "name@example.com")
                    .autocapitalization(.none)
                    
                    InputView(text: $password,
                              title: "Password",
                              placeholder: "Enter your password",
                              isSecureField: true)
                }
                .padding(.horizontal)
                .padding(.top, 32)
                
                // login button
                Button {
                    Task {
                        try await viewModel.login(
                            withEmail: email,
                            password: password
                        )
                    }
                } label: {
                    HStack {
                        Text("Login")
                            .fontWeight(.semibold)
                    }
                    .foregroundColor(.white)
                    .frame(width: UIScreen.main.bounds.width - 32, height: 48)
                }
                .background(Color(.systemPurple))
                .cornerRadius(10)
                .padding(.top, 24)
                
                Spacer()
                
                // sign up link
                NavigationLink {
                    RegistrationView()
                        .navigationBarBackButtonHidden(true)
                } label: {
                    HStack(spacing: 2) {
                        Text("New user?")
                            .foregroundColor(Color(.systemGray))
                            .padding(.trailing, 4)
                        Text("Sign up")
                            .fontWeight(.bold)
                            .foregroundColor(Color(.systemPurple))
                            .padding(.leading, 4)
                    }
                    .font(.system(size: 20))
                }
            }
        }
    }
}

struct LoginView_Previews: PreviewProvider {
    static var previews: some View {
        LoginView()
    }
}
