//
//  TimePeriod.swift
//  Vinance
//
//  Created by Vincent Deli on 13/04/25.
//

import Foundation

public enum TimePeriod: String, CaseIterable, Identifiable {
    case day1 = "1D"
    case week1 = "1W"
    case month1 = "1M"
    case month3 = "3M"
    case yearToDate = "YTD"
    
    public var id: String { self.rawValue }
}
