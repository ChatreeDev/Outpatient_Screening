import { OutpatientScreeningsInterface } from "./IOutpatientScreenings";
import { EmployeesInterface } from "./IEmployee";

export interface HistorySheetsInterface {
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
    Employee: EmployeesInterface;
}
