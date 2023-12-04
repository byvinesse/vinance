//
//  AuthViewModel.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import Foundation

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
        request.httpBody = try! JSONEncoder().encode(LoginRequest(email: email, password: password))
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        let task = URLSession.shared.dataTask(with: request) { data, _, error in
            guard let data = data, error == nil else {
                return
            }

            do {
                let decoder = JSONDecoder()
                decoder.keyDecodingStrategy = .convertFromSnakeCase

                let response = try decoder.decode(LoginResponse.self, from: data)
                print("access_token: \(response.data.accessToken)")
                print("expires_at: \(response.data.expiresAt)")
            } catch {
                print(error)
            }
        }
        task.resume()
    }
    
    func register(withEmail email: String, password: String, username: String) async throws {
        print("Register..")
    }
}


