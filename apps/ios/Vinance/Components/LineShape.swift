//
//  LineShape.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import SwiftUI

struct LineShape: Shape {
    let points: [CGPoint]
    
    func path(in rect: CGRect) -> Path {
        var path = Path()
        
        guard !points.isEmpty else { return path }
        
        let points = points.map { CGPoint(x: $0.x * rect.width, y: $0.y * rect.height) }
        
        path.move(to: points[0])
        
        if points.count == 2 {
            path.addLine(to: points[1])
        } else if points.count > 2 {
            for i in 1..<points.count {
                path.addLine(to: points[i])
            }
        }
        
        return path
    }
}
