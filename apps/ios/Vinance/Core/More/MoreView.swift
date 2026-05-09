//
//  MoreView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

struct MoreView: View {
    var body: some View {
        NavigationStack {
            Text("More")
                .font(.system(size: 24))
                .fontWeight(.bold)
                .foregroundColor(Color(hex: 0x393E46))
        }
    }
}

struct MoreView_Previews: PreviewProvider {
    static var previews: some View {
        MoreView()
    }
}
