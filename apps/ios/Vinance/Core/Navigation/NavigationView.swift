//
//  NavigationView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

struct NavigationView: View {
    var body: some View {
        TabView {
            NavigationStack {
                DashboardView()
            }
            .tabItem {
                Label("Dashboard", systemImage: "house.fill")
            }
            
            NavigationStack {
                PlanningView()
            }
            .tabItem {
                Label("Planning", systemImage: "calendar")
            }
            
            NavigationStack {
                AddNewView()
            }
            .tabItem {
                Label("", systemImage: "plus.circle.fill")
            }
            
            NavigationStack {
                StatisticsView()
            }
            .tabItem {
                Label("Statistics", systemImage: "chart.bar.fill")
            }
            
            NavigationStack {
                MoreView()
            }
            .tabItem {
                Label("More", systemImage: "ellipsis")
            }
        }
        .accentColor(Color(hex: 0x393E46))
    }
}

struct NavigationView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView()
    }
}
