//
//  LoginResponse.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import Foundation

struct LoginResponse: Codable {
    let code: Int32
    let status: String
    let data: LoginData
}

struct LoginData: Codable {
    let accessToken: String
    let expiresAt: String
}
