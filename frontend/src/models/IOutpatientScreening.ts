import { NurseInterface } from "./INurse";
import { HistorySheetInterface } from "./IHistorySheet";
import { EmergencyLevelInterface } from "./IEmergencyLevel";
import { HighBloodPressureLevelsInterface } from "./IHighBloodPressureLevel";
import { DiabetesLevelInterface } from "./IDiabetesLevel";
import { ObesityLevelInterface } from "./IObesityLevel";



export interface OutpatientScreeningsInterface {
    ID: number;
    Time: Date;

    NurseID: number;
    Nurse: NurseInterface;

    HistorySheetID: number;
    HistorySheet: HistorySheetInterface;

    EmergencyLevelID: number;
    EmergencyLevel: EmergencyLevelInterface;

    HighBloodPressureLevelID: number;
    HighBloodPressureLevel: HighBloodPressureLevelsInterface;

    DiabetesLevelID: number;
    DiabetesLevel: DiabetesLevelInterface;

    ObesityLevelID: number;
    ObesityLevel: ObesityLevelInterface;

    Note: string;
}

