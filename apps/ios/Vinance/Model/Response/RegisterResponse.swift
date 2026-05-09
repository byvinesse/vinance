//
//  RegisterResponse.swift
//  Vinance
//
//  Created by Vincent Deli on 14/01/24.
//

import Foundation

struct RegisterResponse: Codable {
    let code: Int32
    let status: String
    let data: Bool
}
