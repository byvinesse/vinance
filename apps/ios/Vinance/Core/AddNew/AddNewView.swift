//
//  AddNewView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

struct AddNewView: View {
    var body: some View {
        NavigationStack {
            Text("Add New")
                .font(.system(size: 24))
                .fontWeight(.bold)
                .foregroundColor(Color(hex: 0x393E46))
        }
    }
}

struct AddNewView_Previews: PreviewProvider {
    static var previews: some View {
        AddNewView()
    }
}
