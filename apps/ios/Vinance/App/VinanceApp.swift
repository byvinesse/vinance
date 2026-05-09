//
//  VinanceApp.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

@main
struct VinanceApp: App {
    @StateObject var viewModel = AuthViewModel()
    
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(viewModel)
        }
    }
}
