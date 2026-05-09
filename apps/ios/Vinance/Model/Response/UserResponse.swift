//
//  UserResponse.swift
//  Vinance
//
//  Created by Vincent Deli on 06/04/25.
//

import Foundation

struct UserResponse: Codable {
    let code: Int32
    let status: String
    let data: UserData
}

struct UserData: Codable {
    let email: String
    let username: String
    let phoneNumber: String
    let gender: String
    let dateOfBirth: Int64
}
