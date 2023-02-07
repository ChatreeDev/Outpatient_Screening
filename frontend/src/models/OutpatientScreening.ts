import { HistorySheetsInterface } from "./IHistorySheet";
import { HighBloodPressureLevelsInterface } from "./IHighBloodPressureLevel";
import { DiabetesLevelsInterface } from "./IDiabetesLevel";

/*  Interface มีเอาไว้ทำอะไร ? Ans. เป็นการกำหนดโครงสร้างของข้อมูล
    เวลาเราเรียกใช้มันจะได้หาชื่อได้ง่าย
    key ลัด : command + . => ทำ quick fiq แล้วก็กด auto import
*/

export interface OutpatientScreeningsInterface {
    ID: number;

    FoodTime: Date;

    HighBloodPressureLevelID: number;
    HighBloodPressureLevel: HighBloodPressureLevelsInterface;

    BookingID: number;
    Booking: HistorySheetsInterface;

    OutpatientScreeningDiabetesLevel: OutpatientScreeningDiabetesLevelInterface[];

    TotalPrice: number;
}

// เอาสองตารางนี้ไว้ด้วยกันเพราะมันเป็นตารางบันทึกทั้งคู่
export interface OutpatientScreeningDiabetesLevelInterface {
    ID: number;

    OutpatientScreeningID: number;
    OutpatientScreening: OutpatientScreeningsInterface;

    FoodSetID: number;
    FoodSet: DiabetesLevelsInterface;

    Quantity: number;
}