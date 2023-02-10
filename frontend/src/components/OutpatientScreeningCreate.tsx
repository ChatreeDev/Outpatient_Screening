import React, { useEffect } from "react";
import { Link as RouterLink, useParams } from "react-router-dom";
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import { CardActionArea, MenuItem } from '@mui/material';
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";

import { NurseInterface } from "../models/INurse";
import { HistorySheetInterface } from "../models/IHistorySheet";
import { EmergencyLevelInterface } from "../models/IEmergencyLevel";
import { HighBloodPressureLevelsInterface } from "../models/IHighBloodPressureLevel";
import { DiabetesLevelInterface } from "../models/IDiabetesLevel";
import { ObesityLevelInterface } from "../models/IObesityLevel";
import { OutpatientScreeningsInterface } from "../models/IOutpatientScreening";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export default function OutpatienScreeningsCreate() {
  // ดึง HistorySheetId มาจาก url ตอนกดเลือกใบการซักประวัติ ex. "localhost:3000/OutpatienScreeningCreate/2" จะได้ HistorySheetId = 2
  const { HistorySheetId } = useParams();

  const [HistorySheet, setHistorySheet] = React.useState<HistorySheetInterface>();
  const [EmergencyLevel, setEmergencyLevel] = React.useState<EmergencyLevelInterface | null>();
  const [HighBloodPressureLevels, setHighBloodPressureLevels] = React.useState<HighBloodPressureLevelsInterface>();
  const [DiabetesLevel, setDiabetesLevel] = React.useState<DiabetesLevelInterface>();
  const [ObesityLevel, setObesityLevel] = React.useState<ObesityLevelInterface>();
  
  const [OutpatienScreenings, setOutpatienScreenings] = React.useState<Partial<OutpatientScreeningsInterface>>({
    HistorySheetID: Number(HistorySheetId), EmergencyLevelID: Number(EmergencyLevel),HighBloodPressureLevelID: Number(HighBloodPressureLevels),DiabetesLevelID: Number(DiabetesLevel),ObesityLevelID: Number(ObesityLevel), Time: new Date()       
  });

  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  //Get Function
  const getHistorySheet = async () => {
    const apiUrl = `http://localhost:8080/HistorySheet/${HistorySheetId}`;       
    const requestOptions = {
      method: "GET",                                       
      headers: {
        "Content-Type": "application/json",                
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }

    fetch(apiUrl, requestOptions)                          
      .then(response => response.json())                   
      .then(res => {
        console.log(res);                                  
        if (res.data) {                                    
          setHistorySheet(res.data);                           
        } else {                                           
          console.log(res.error);
        }
      })
  }
  const getEmergencyLevel = async () => {
    const apiUrl = "http://localhost:8080/foodpayment_types";
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(res => {
        console.log(res);
        if (res.data) {
          setEmergencyLevel(res.data);
        } else {
          console.log(res.error);
        }
      })
  }

  const getHighBloodPressureLevels = async () => {                       
    const apiUrl = "http://localhost:8080/HighBloodPressureLevels";
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(res => {
        console.log(res);
        if (res.data) {
          setHighBloodPressureLevels(res.data);
        } else {
          console.log(res.error);
        }
      })
  }
  const getDiabetesLevel = async () => {                       
    const apiUrl = "http://localhost:8080/DiabetesLevel";
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(res => {
        console.log(res);
        if (res.data) {
          setDiabetesLevel(res.data);
        } else {
          console.log(res.error);
        }
      })
  }
  const getObesityLevel = async () => {                       
    const apiUrl = "http://localhost:8080/ObesityLevel";
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(res => {
        console.log(res);
        if (res.data) {
          setObesityLevel(res.data);
        } else {
          console.log(res.error);
        }
      })
  }
  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const submit = () => {
    let data = {
      HistorySheetID: convertType(OutpatienScreenings.HistorySheetID),
      EmergencyLevelID: convertType(OutpatienScreenings. EmergencyLevelID),
      HighBloodPressureLevelID: convertType(OutpatienScreenings.HighBloodPressureLevelID),
      DiabetesLevelID: convertType(OutpatienScreenings.DiabetesLevelID),
      ObesityLevelID: convertType(OutpatienScreenings.ObesityLevelID),
      //Time: selectedDate,
      //Note: OutpatienScreenings.Note ?? "",
    };
    console.log(data)

    useEffect(() => {
      getHistorySheet();
      getEmergencyLevel();
      getHighBloodPressureLevels();
      getDiabetesLevel();
      getObesityLevel();
    }, []);

    const apiUrl = "http://localhost:8080/OutpatienScreenings";
    const requestOptionsPost = {
      
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/OutpatienScreenings`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          //setErrorMessage("")
          //ClearForm();
        } else {
          console.log("บันทึกไม่ได้")
          setSuccess(true);
          //setErrorMessage("")
        }
  });

  return (
   //Create OutpatienScreenings Form Using MUI Autocomplete
   
    <div>
      <Container maxWidth="md">
      <h1>OutpatienScreenings</h1>
      <form className={""} noValidate autoComplete="off"></form>
        <div>
          <TextField
            id="HistorySheetID"
            label="HistorySheetID"
            variant="outlined"
            value={OutpatienScreenings.HistorySheetID}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                //HistorySheetID: e.target.value, 
              })
            }
          />
        </div>
        <div>
          <TextField
            id="EmergencyLevelID"
            label="EmergencyLevelID"
            variant="outlined"
            value={OutpatienScreenings.EmergencyLevelID}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                //EmergencyLevelID: e.target.value,
              })
            }
          />
        </div>
        <div>
          <TextField
            id="HighBloodPressureLevelID"
            label="HighBloodPressureLevelID"
            variant="outlined"
            value={OutpatienScreenings.HighBloodPressureLevelID}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                //HighBloodPressureLevelID: e.target.value,
              })
            }
          />
        </div>
        <div>
          <TextField
            id="DiabetesLevelID"
            label="DiabetesLevelID"
            variant="outlined"
            value={OutpatienScreenings.DiabetesLevelID}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
               //DiabetesLevelID: e.target.value,
              })
            }
          />
        </div>
        <div>
          <TextField
            id="ObesityLevelID"
            label="ObesityLevelID"
            variant="outlined"
            value={OutpatienScreenings.ObesityLevelID}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                //ObesityLevelID: e.target.value,
              })
            }
          />
        </div>
        <div>
          <TextField
            id="Time"
            label="Time"
            variant="outlined"
            value={OutpatienScreenings.Time}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                //Time: e.target.value,
              })
            }
          />  
        </div>
        <div>
          <TextField
            id="Note"
            label="Note"
            variant="outlined"
            value={OutpatienScreenings.Note}
            onChange={(e) =>
              setOutpatienScreenings({
                ...OutpatienScreenings,
                Note: e.target.value,
              })
            }
          />
        </div>
        <div>
          <Button
            variant="contained"
            color="primary"
            onClick={() => submit()}
          > 
            Submit
          </Button>
        </div>
      </Container>
    </div>
  );
}
}
//</form>