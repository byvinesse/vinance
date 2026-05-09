//
//  ContentView.swift
//  Vinance
//
//  Created by Vincent Deli on 03/12/23.
//

import SwiftUI

struct ContentView: View {
    @EnvironmentObject var viewModel: AuthViewModel

    var body: some View {
        // VNC-36: Bypass auth wall — show home screen with mock data for first iteration.
        // Restore the auth check below once backend integration is complete:
        //
        //   if viewModel.userSession != nil {
        //       AppTabView()
        //   } else {
        //       LoginView()
        //   }
        //
        AppTabView()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
            .environmentObject(AuthViewModel())
    }
}
