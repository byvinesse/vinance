//
//  AuthViewModel.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import Foundation
import SwiftUI

// To use the networking components, you'll need to add these imports:
// The actual import statements will depend on your project structure.
// Here are a few options:

// If network components are in the same module/target:
// No additional imports needed

// If network components are in a separate module/target named "Network":
// import Network 

// If you're using Swift Package Manager with a package named "NetworkKit":
// import NetworkKit

enum AuthError: LocalizedError {
    case invalidEmail
    case invalidPassword
    case passwordMismatch
    case invalidURL
    case networkError(Error)
    case decodingError(Error)
    case unknown
    
    var errorDescription: String? {
        switch self {
        case .invalidEmail:
            return "Please enter a valid email address."
        case .invalidPassword:
            return "Password must be at least 8 characters long."
        case .passwordMismatch:
            return "Passwords do not match."
        case .invalidURL:
            return "Invalid URL."
        case .networkError(let error):
            return "Network error: \(error.localizedDescription)"
        case .decodingError(let error):
            return "Failed to process response: \(error.localizedDescription)"
        case .unknown:
            return "An unknown error occurred."
        }
    }
}

@MainActor
class AuthViewModel: ObservableObject {
    @Published var userSession: String?
    @Published var currentUser: User?
    
    private let userDefaults = UserDefaults.standard
    private let userSessionKey = "userSession"
    private let networkService: NetworkServiceProtocol
    
    init(networkService: NetworkServiceProtocol = NetworkService()) {
        self.networkService = networkService
        self.userSession = userDefaults.string(forKey: userSessionKey)
        if userSession != nil {
            Task {
                try? await fetchCurrentUser()
            }
        }
    }
    
    func login(withEmail email: String, password: String) async throws {
        do {
            let loginRequest = LoginRequest(email: email, password: password)
            let response: LoginResponse = try await networkService.request(
                endpoint: "/auth/v1/login",
                method: .post,
                body: loginRequest,
                headers: nil
            )
            
            self.userSession = response.data.accessToken
            userDefaults.set(self.userSession, forKey: userSessionKey)
            try await fetchCurrentUser()
        } catch let error as NetworkError {
            switch error {
            case .invalidURL:
                throw AuthError.invalidURL
            case .decodingFailed(let decodingError):
                throw AuthError.decodingError(decodingError)
            case .requestFailed, .unknown:
                throw AuthError.networkError(error)
            }
        } catch {
            throw AuthError.unknown
        }
    }
    
    func register(withEmail email: String, password: String, confirmPassword: String) async throws -> Bool {
        do {
            try validateRegisterInput(email: email, password: password, confirmPassword: confirmPassword)
            
            let registerRequest = RegisterRequest(email: email, password: password)
            let response: RegisterResponse = try await networkService.request(
                endpoint: "/auth/v1/register",
                method: .post,
                body: registerRequest,
                headers: nil
            )
            
            return response.data
        } catch let error as NetworkError {
            switch error {
            case .invalidURL:
                throw AuthError.invalidURL
            case .decodingFailed(let decodingError):
                throw AuthError.decodingError(decodingError)
            case .requestFailed, .unknown:
                throw AuthError.networkError(error)
            }
        } catch let error as AuthError {
            throw error
        } catch {
            throw AuthError.unknown
        }
    }
    
    func signOut() {
        self.userSession = nil
        self.currentUser = nil
        userDefaults.removeObject(forKey: userSessionKey)
    }
    
    private func fetchCurrentUser() async throws {
        guard let token = userSession else { 
            throw AuthError.unknown
        }
        
        // Example of how to use the authenticated request
        // Implement this based on your API
        // let user: UserResponse = try await networkService.requestWithAuth(
        //    endpoint: "/auth/v1/user",
        //    method: .get,
        //    body: nil,
        //    token: token
        // )
        // self.currentUser = user.data
    }
    
    private func validateRegisterInput(email: String, password: String, confirmPassword: String) throws {
        let emailRegex = "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}"
        guard !email.isEmpty, email.range(of: emailRegex, options: .regularExpression) != nil else {
            throw AuthError.invalidEmail
        }
        
        guard !password.isEmpty, password.count >= 8 else {
            throw AuthError.invalidPassword
        }
        
        guard confirmPassword == password else {
            throw AuthError.passwordMismatch
        }
    }
}


