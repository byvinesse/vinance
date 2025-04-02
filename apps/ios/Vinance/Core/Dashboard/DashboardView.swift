//
//  DashboardView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

struct DashboardView: View {
    var body: some View {
        NavigationStack {
            Text("Vinance Dashboard")
                .font(.system(size: 24))
                .fontWeight(.bold)
                .foregroundColor(Color(hex: 0x393E46))
        }
    }
}

struct DashboardView_Previews: PreviewProvider {
    static var previews: some View {
        DashboardView()
    }
}
