//
//  ContentView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        LoginView()
            .tabItem() {
                Image(systemName: "person.crop.circle.badge.checkmark")
                Text("Login")
            }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
