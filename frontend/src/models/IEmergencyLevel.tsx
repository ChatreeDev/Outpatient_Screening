import { OutpatientScreeningsInterface } from "./IOutpatientScreening";

export interface EmergencyLevelInterface {
    id: number;
    Level : string;
    AssessmentForms : string;
    HistoryTaking : string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}