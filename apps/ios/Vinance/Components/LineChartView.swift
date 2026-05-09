//
//  LinechartView.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct LineChartView: View {
    let data: [ChartPoint]
    let showUnderBudget: Bool
    var underBudgetValue: String = ""
    let colors: [Color]
    
    private var minValue: Double {
        data.map { $0.value }.min() ?? 0
    }
    
    private var maxValue: Double {
        data.map { $0.value }.max() ?? 0
    }
    
    private var range: Double {
        maxValue - minValue
    }
    
    private func normalize(_ value: Double) -> Double {
        if range == 0 { return 0.5 }
        return (value - minValue) / range
    }
    
    var body: some View {
        ZStack(alignment: .center) {
            // Chart line
            LineShape(points: data.enumerated().map { index, point in
                CGPoint(
                    x: CGFloat(index) / CGFloat(data.count - 1),
                    y: 1.0 - CGFloat(normalize(point.value))
                )
            })
            .trim(from: 0, to: 1)
            .stroke(
                LinearGradient(
                    gradient: Gradient(colors: colors),
                    startPoint: .leading,
                    endPoint: .trailing
                ),
                style: StrokeStyle(lineWidth: 2, lineCap: .round, lineJoin: .round)
            )
            
            // Points on the line
            ForEach(data.indices, id: \.self) { index in
                let point = data[index]
                let position = CGPoint(
                    x: CGFloat(index) / CGFloat(data.count - 1),
                    y: 1.0 - CGFloat(normalize(point.value))
                )
                
                // Only show point for the last value
                if index == data.count - 1 {
                    Circle()
                        .fill(colors.last ?? .green)
                        .frame(width: 8, height: 8)
                        .position(
                            x: position.x * UIScreen.main.bounds.width * 0.9,
                            y: position.y * 60
                        )
                    
                    // Under budget indicator
                    if showUnderBudget {
                        Text("$\(underBudgetValue) under")
                            .font(.system(size: 10, weight: .bold))
                            .padding(4)
                            .background(Color.green)
                            .foregroundColor(.white)
                            .cornerRadius(4)
                            .position(
                                x: position.x * UIScreen.main.bounds.width * 0.9,
                                y: max(position.y * 60 - 20, 10)
                            )
                    }
                }
            }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
    }
}
