import { Container } from '@mui/system'
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';


function Home() {
  return (
    <div>
      <Container maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบการคัดกรองคนไข้นอก</h1>
        <h4><u>Requirements</u></h4>
        <p>
        ระบบคัดกรองคนไข้นอกเป็นระบบสำหรับการประเมินคนไข้นอกเพื่อค้นหาปัญหาและความเสี่ยงโดยการจัดลำดับความเร่งด่วนในการตรวจรักษาของคนไข้นอก
         โดยพยาบาลซึ่งเป็นผู้ใช้ระบบจะทำการประเมินอาการเสี่ยงของคนไข้นอก โดยจะแบ่งตามระดับความรุนแรงซึ่งการประเมินจะถูกแบ่งออกเป็น 3 ประเภทได้แก่
          Emergency (ภาวะฉุกเฉิน), Urgent (ภาวะเร่งด่วน), Non-urgent (ภาวะไม่เร่งด่วน) พยาบาลจะทำการประเมินอาการเสี่ยงของคนไข้นอกด้วยตนเอง 
          อีกทั้งยังทำการคัดกรองโดยการประเมินคนไข้ที่เสี่ยงเบาหวาน, ประเมินคนไข้ที่เสี่ยงความดันโลหิตสูง, ประเมินคนไข้ที่เสี่ยงเป็นโรคอ้วน เป็นต้น
        </p>
        <br />
        <Card sx={{ maxWidth: 850 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image='https://images.unsplash.com/photo-1588543385566-413e13a51a24?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80'
              //image="https://images.unsplash.com/photo-1529148482759-b35b25c5f217?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80"
            />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Emergency
              </Typography>
              <Typography variant="body2" color="text.secondary">
                ยินดีต้อนรับพยาบาลทุกท่าน พยาบาลสามารถคัดกรองคนไข้นอกได้อย่างถี่ถ้วนจากระบบนี้
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
        <Card sx={{ maxWidth: 850 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image='https://images.unsplash.com/photo-1603398938378-e54eab446dde?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80'
              //image="https://images.unsplash.com/photo-1589010588553-46e8e7c21788?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2160&q=80"
            />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                HighBloodPressure, Diabetes, Obesity
              </Typography>
              <Typography variant="body2" color="text.secondary">
                พยาบาลทุกท่านสามารถคัดกรองคนไข้นอกได้ที่หน้าถัดไป
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      </Container>
    </div>
  )
}

export default Home