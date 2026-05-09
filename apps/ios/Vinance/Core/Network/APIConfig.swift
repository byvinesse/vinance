//
//  APIConfig.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import Foundation

enum APIEnvironment {
    case development
    case staging
    case production
    
    var baseURL: String {
        switch self {
        case .development:
            return "http://localhost:8080"
        case .staging:
            return "https://staging-api.vinance.app"
        case .production:
            return "https://api.vinance.app"
        }
    }
}

struct APIConfig {
    #if DEBUG
    static let environment: APIEnvironment = .development
    #else
    static let environment: APIEnvironment = .production
    #endif
    
    static var baseURL: String {
        environment.baseURL
    }
    
    static func endpoint(_ path: String) -> URL? {
        URL(string: "\(baseURL)\(path)")
    }
}
