//
//  NetworkService.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import Foundation

enum HTTPMethod: String {
    case get = "GET"
    case post = "POST"
    case put = "PUT"
    case patch = "PATCH"
    case delete = "DELETE"
}

enum NetworkError: Error {
    case invalidURL
    case requestFailed(statusCode: Int)
    case decodingFailed(Error)
    case unknown(Error)
}

protocol NetworkServiceProtocol {
    func request<T: Decodable>(endpoint: String, method: HTTPMethod, body: Encodable?, headers: [String: String]?) async throws -> T
    func requestWithAuth<T: Decodable>(endpoint: String, method: HTTPMethod, body: Encodable?, token: String) async throws -> T
}

class NetworkService: NetworkServiceProtocol {
    private let timeoutInterval: TimeInterval
    
    init(timeoutInterval: TimeInterval = 30.0) {
        self.timeoutInterval = timeoutInterval
    }
    
    func request<T: Decodable>(endpoint: String, method: HTTPMethod, body: Encodable? = nil, headers: [String: String]? = nil) async throws -> T {
        guard let url = APIConfig.endpoint(endpoint) else {
            throw NetworkError.invalidURL
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = method.rawValue
        request.timeoutInterval = timeoutInterval
        
        // Set default headers
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        
        // Add custom headers if provided
        headers?.forEach { key, value in
            request.setValue(value, forHTTPHeaderField: key)
        }
        
        if let body = body {
            let encoder = JSONEncoder()
            request.httpBody = try encoder.encode(body)
        }
        
        let (data, response) = try await URLSession.shared.data(for: request)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NetworkError.unknown(NSError(domain: "Invalid response", code: 0, userInfo: nil))
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            throw NetworkError.requestFailed(statusCode: httpResponse.statusCode)
        }
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        
        do {
            return try decoder.decode(T.self, from: data)
        } catch {
            throw NetworkError.decodingFailed(error)
        }
    }
    
    func requestWithAuth<T: Decodable>(endpoint: String, method: HTTPMethod, body: Encodable? = nil, token: String) async throws -> T {
        let headers = ["Authorization": "Bearer \(token)"]
        return try await request(endpoint: endpoint, method: method, body: body, headers: headers)
    }
}
