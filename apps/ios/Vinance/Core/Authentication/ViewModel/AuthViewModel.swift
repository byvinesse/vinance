//
//  AuthViewModel.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import Foundation

@MainActor
class AuthViewModel: ObservableObject {
    @Published var userSession: String?
    @Published var currentUser: User?
    
    init() {
        
    }
    
    func login(withEmail email: String, password: String) async throws {
        guard let url = URL(string: "http://localhost:8080/v1/login") else {
            throw CustomError.invalidURL
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        
        do {
            let requestBody = try JSONEncoder().encode(LoginRequest(email: email, password: password))
            request.httpBody = requestBody
            request.setValue("application/json", forHTTPHeaderField: "Content-Type")
            
            let (data, _) = try await URLSession.shared.data(for: request)
            
            let decoder = JSONDecoder()
            decoder.keyDecodingStrategy = .convertFromSnakeCase
            
            let response = try decoder.decode(LoginResponse.self, from: data)
            
            self.userSession = response.data.accessToken
        } catch {
            print(error)
        }
    }
    
    func register(withEmail email: String, password: String, confirmPassword: String) async throws -> Bool {
        if !validateRegisterPayload(email: email, password: password, confirmPassword: confirmPassword) {
            throw CustomError.invalidData
        }
        
        guard let url = URL(string: "http://localhost:8080/v1/register") else {
            throw CustomError.invalidURL
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        
        do {
            let requestBody = try JSONEncoder().encode(RegisterRequest(email: email, password: password))
            request.httpBody = requestBody
            request.setValue("application/json", forHTTPHeaderField: "Content-Type")
            
            let (data, _) = try await URLSession.shared.data(for: request)
            
            let decoder = JSONDecoder()
            decoder.keyDecodingStrategy = .convertFromSnakeCase
            
            let response = try decoder.decode(RegisterResponse.self, from: data)
            return response.data
        } catch {
            print(error)
        }
        
        print("Error: Registration failed.")
        return false
    }
    
    func validateRegisterPayload(email: String, password: String, confirmPassword: String) -> Bool {
        let emailRegex = "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}"
        guard email.range(of: emailRegex, options: .regularExpression) != nil, !email.isEmpty else {
            print("Error: Invalid email format.")
            return false
        }
        
        guard !password.isEmpty, password.count >= 8 else {
            print("Error: Invalid password format.")
            return false
        }
        
        guard confirmPassword == password else {
            print("Error: Password does not match.")
            return false
        }
        
        return true
    }
}


