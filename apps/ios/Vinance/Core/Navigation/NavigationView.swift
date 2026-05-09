//
//  NavigationView.swift
//  Vinance
//
//  Created by Vincent Deli on 02/04/25.
//

import SwiftUI

/// Root tab-bar container.
/// Tab structure mirrors the design reference (VNC-36):
///   Home · Txns · Analytics · Budgets  +  FAB overlay
struct AppTabView: View {

    init() {
        configureTabBarAppearance()
    }

    var body: some View {
        ZStack(alignment: .bottom) {
            TabView {
                // 1 — Home
                NavigationStack {
                    HomeView()
                }
                .tabItem {
                    Label("Home", systemImage: "house.fill")
                }

                // 2 — Transactions (placeholder — full screen in future ticket)
                NavigationStack {
                    PlanningView()
                }
                .tabItem {
                    Label("Txns", systemImage: "list.bullet.rectangle.portrait")
                }

                // 3 — Analytics (placeholder — full screen in future ticket)
                NavigationStack {
                    StatisticsView()
                }
                .tabItem {
                    Label("Analytics", systemImage: "chart.bar.fill")
                }

                // 4 — Budgets (placeholder — full screen in future ticket)
                NavigationStack {
                    MoreView()
                }
                .tabItem {
                    Label("Budgets", systemImage: "target")
                }
            }
            .accentColor(Color.vAccent2)
        }
    }

    // MARK: - Tab Bar Styling

    private func configureTabBarAppearance() {
        let appearance = UITabBarAppearance()
        appearance.configureWithOpaqueBackground()
        appearance.backgroundColor = UIColor(Color.vSurface).withAlphaComponent(0.95)

        // Normal item style
        let normalAttrs: [NSAttributedString.Key: Any] = [
            .foregroundColor: UIColor(Color.vText3)
        ]
        // Selected item style
        let selectedAttrs: [NSAttributedString.Key: Any] = [
            .foregroundColor: UIColor(Color.vAccent2)
        ]

        appearance.stackedLayoutAppearance.normal.titleTextAttributes   = normalAttrs
        appearance.stackedLayoutAppearance.selected.titleTextAttributes = selectedAttrs
        appearance.stackedLayoutAppearance.normal.iconColor   = UIColor(Color.vText3)
        appearance.stackedLayoutAppearance.selected.iconColor = UIColor(Color.vAccent2)

        UITabBar.appearance().standardAppearance   = appearance
        UITabBar.appearance().scrollEdgeAppearance = appearance
        UITabBar.appearance().tintColor            = UIColor(Color.vAccent2)
    }
}

struct AppTabView_Previews: PreviewProvider {
    static var previews: some View {
        AppTabView()
            .preferredColorScheme(.dark)
    }
}
