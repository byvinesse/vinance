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
    
    func register(withEmail email: String, password: String, username: String) async throws {
        print("Register..")
    }
}


