//
//  ProfileView.swift
//  Vinance
//
//  Created by Vincent Deli on 04/12/23.
//

import SwiftUI

struct ProfileView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    
    var body: some View {
        let currentUser = authViewModel.currentUser
        
        List {
            Section {
                HStack {
                    Text(currentUser?.initials ?? "UU")
                        .font(.title)
                        .fontWeight(.semibold)
                        .foregroundColor(.white)
                        .frame(width: 72, height: 72)
                        .background(Color(.systemGray3))
                    .clipShape(Circle())
                    
                    VStack(alignment: .leading, spacing: 4) {
                        Text(currentUser?.username ?? "User")
                            .font(.subheadline)
                            .fontWeight(.semibold)
                            .padding(.top, 4)
                        
                        Text(currentUser?.email ?? "Email")
                            .font(.footnote)
                            .foregroundColor(.gray)
                    }
                }
            }
            
//            Section("General") {
//                HStack {
//                    SettingsRowView(imageName: "gear", title: "Version", tintColor: Color(.systemGray))
//
//                    Spacer()
//
//                    Text("1.0.0")
//                        .font(.subheadline)
//                        .foregroundColor(.gray)
//                }
//            }
            
            Section("General") {
                Button {
                    Task {
                        authViewModel.signOut()
                    }
                } label: {
                    SettingsRowView(imageName: "arrow.left.circle.fill",
                                    title: "Sign Out",
                                    tintColor: .red)
                }
                
//                Button {
//                    print("Delete account..")
//                } label: {
//                    SettingsRowView(imageName: "xmark.circle.fill",
//                                    title: "Delete Account",
//                                    tintColor: .red)
//                }
            }
        }
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
