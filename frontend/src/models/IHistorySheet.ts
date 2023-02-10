import { OutpatientScreeningsInterface } from "./IOutpatientScreening";
import { NurseInterface } from "./INurse";

export interface HistorySheetInterface {
    ID: number;
    Weight: number;
    Height: number;
    BMI: number;
    Temperature: number;
    SystolicBloodPressure: number;
    DiastolicBloodPressure: number;
    HeartRate: number;
    RespiratoryRate: number;
    OxygenSaturation: number;
    DrugAllergy: string;
    PatientSymptoms: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
    Nurse: NurseInterface[];
}
