//
//  RegistrationView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

struct RegistrationView: View {
    @State private var email = ""
    @State private var username = ""
    @State private var password = ""
    @State private var confirmPassword = ""
    @State private var registrationSuccess = false
    @Environment(\.dismiss) var dismiss
    @EnvironmentObject var viewModel: AuthViewModel
    
    var body: some View {
        VStack {
            // image
            Image("vinance-logo-bgless")
                .resizable()
                .scaledToFill()
                .frame(width: 344, height: 275)
            
            // hero text
            VStack(alignment: .leading) {
                Text("Hello")
                    .font(.system(size: 48))
                    .fontWeight(.bold)
                    .foregroundColor(Color(hex: 0x393E46))
                    
                Text("Begin your journey here")
                    .font(.system(size: 24))
                    .fontWeight(.bold)
                    .foregroundColor(Color(hex: 0x929AAB))
            }
            .padding(.top, -32)
            .padding(.bottom, 32)
            
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
                
                InputView(text: $confirmPassword,
                          title: "Confirm Password",
                          placeholder: "Confirm your password",
                          isSecureField: true)
                
                // register button
                Button {
                    Task {
                        do {
                            let success = try await viewModel.register(
                                withEmail: email,
                                password: password,
                                confirmPassword: confirmPassword
                            )
                            
                            registrationSuccess = success
                        } catch {
                            print("Error: Registration failed.")
                        }
                    }
                } label: {
                    HStack {
                        Text("Register")
                            .fontWeight(.semibold)
                    }
                    .foregroundColor(.white)
                    .frame(width: UIScreen.main.bounds.width - 32, height: 48)
                }
                .background(Color(.systemPurple))
                .cornerRadius(10)
                .padding(.top, 24)
                .fullScreenCover(isPresented: $registrationSuccess) {
                    LoginView()
                }
                
                Spacer()
                
                Button {
                    dismiss()
                } label: {
                    HStack(spacing: 2) {
                        Text("Have an account?")
                            .foregroundColor(Color(.systemGray))
                            .padding(.trailing, 4)
                        Text("Sign in")
                            .fontWeight(.bold)
                            .foregroundColor(Color(.systemPurple))
                            .padding(.leading, 4)
                    }
                    .font(.system(size: 20))
                }
            }
            .padding(.horizontal)
            .padding(.top, 12)
        }
    }
}

struct RegistrationView_Previews: PreviewProvider {
    static var previews: some View {
        RegistrationView()
    }
}
