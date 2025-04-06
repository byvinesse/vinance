//
//  User.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import Foundation

struct User: Codable {
    let username: String
    let email: String
    let phoneNumber: String
    let gender: String
    let dateOfBirth: Int64
    
    var initials: String {
        let formatter = PersonNameComponentsFormatter()
        if let components = formatter.personNameComponents(from: username) {
            formatter.style = .abbreviated
            return formatter.string(from: components)
        }
        
        return ""
    }
}

extension User {
    static var MOCK_USER = User(
        username: "Vincent Deli",
        email: "vincentkdeli@gmail.com",
        phoneNumber: "08123456789",
        gender: "M",
        dateOfBirth: 946684800000
    )
}
