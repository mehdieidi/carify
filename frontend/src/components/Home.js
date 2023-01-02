import "./Home.css";
import Select from "react-select";
import axios from "axios";
import React, { useState, useEffect } from "react";
function Home() {
  const brand = [{ label: "پراید 141 ساده", value: 1 }];
  const [show, setShow] = useState(false);
  const [predict, setPredict] = useState(0);

  const year = [
    { label: "1394", value: 1 },
    { label: "1393", value: 2 },
    { label: "1392", value: 3 },
    { label: "1391", value: 4 },
    { label: "1390", value: 5 },
    { label: "1389", value: 6 },
    { label: "1388", value: 7 },
    { label: "1387", value: 8 },
    { label: "1386", value: 9 },
    { label: "1385", value: 10 },
    { label: "1384", value: 11 },
    { label: "1383", value: 12 },
    { label: "1382", value: 13 },
  ];
  const bime = [
    { label: "1 ماه", value: 1 },
    { label: "2 ماه", value: 2 },
    { label: "3 ماه", value: 3 },
    { label: "4 ماه", value: 4 },
    { label: "5 ماه", value: 5 },
    { label: "6 ماه", value: 6 },
    { label: "7 ماه", value: 7 },
    { label: "8 ماه", value: 8 },
    { label: "9 ماه", value: 9 },
    { label: "10 ماه", value: 10 },
    { label: "11 ماه", value: 11 },
    { label: "12 ماه", value: 12 },
  ];

  const [body, setBody] = useState([]);
  const [color, setColor] = useState([]);
  const [frontChassis, setFrontChassis] = useState([]);
  const [gearbox, setGearbox] = useState([]);
  const [motor, setMotor] = useState([]);
  const [rearChassis, setRearChassis] = useState([]);

  const [body_status_value, setbody_status_value] = useState(0);
  const [color_value, setcolor_value] = useState(0);
  const [front_chassis_status_value, setfront_chassis_status_value] =
    useState(0);
  const [gearbox_, setgearbox] = useState(0);
  const [insurance, setinsurance] = useState(0);
  const [motor_status_value, setmotor_status_value] = useState(0);
  const [rear_chassis_status_value, setrear_chassis_status_value] = useState(0);
  const [usage_km, setusage_km] = useState(0);
  const [year_, setyear] = useState(0);

  useEffect(() => getInfo(), []);
  useEffect(() => getInfo(), []);

  function getInfo() {
    console.log("hee lllo");
    axios
      .get(`http://45.67.85.169:8080/v1/site/settings/get`)
      .then((resp) => {
        console.log("gett");
        console.log(resp.data);
        setBody(resp.data.data.body);
        setColor(resp.data.data.color);
        setFrontChassis(resp.data.data.front_chassis_status);
        setGearbox(resp.data.data.gearbox);
        setMotor(resp.data.data.motor_status);
        setRearChassis(resp.data.data.rear_chassis_status);
        console.log(body);
        console.log(color.map((a) => a.name));
        console.log(frontChassis);
        console.log(gearbox);
        console.log(motor);
        console.log(rearChassis);
      })
      .catch((e) => {
        console.log("failed");
      });
  }
  //
  function submit(e) {
    e.preventDefault();
    console.log("submit func");
    axios
      .post(`http://45.67.85.169:8080/v1/costs/predict`, {
        body_status_value: body_status_value,
        color_value: color_value,
        front_chassis_status_value: front_chassis_status_value,
        gearbox: gearbox_,
        insurance: insurance,
        motor_status_value: motor_status_value,
        rear_chassis_status_value: rear_chassis_status_value,
        usage_km: usage_km,
        year: year_,
      })
      .then(async function (response) {
        console.log(response.data);

        if (response.status === 200) {
          console.log(response.data.data);
          setPredict(response.data.data);
          setShow(true);
        }
      })
      .catch((err) => {
        console.log("error");
        if (err.response.status === 400) {
          console.log("invalid request");
        }
      });
  }

  return (
    <div className="home-page">
      <nav className="nav-bar">
        <ul className="nav-list">
          {/* <li>About</li> */}
          <li>Home</li>
          <li>
            <a
              href="https://github.com/mehdieidi/carify"
              style={{ color: "white", textDecoration: "none" }}
            >
              Github
            </a>
          </li>
        </ul>
      </nav>
      <header className="header">
        <div id="header-des">
          <h1>پیش بینی قیمت خودرو با استفاده از خودروهای بازار ایران</h1>
          <p style={{ marginTop: "10vh" }}>
            با جمع آوری اطلاعات خودرو از سایت دیوار، قیمت خودرو را پیش بینی می
            کند
          </p>
          {/* <button>شروع</button> */}
        </div>
        <div id="header-photo">
          <img src={require("./RTX2d5Q7X.jpg")} alt="cars" />
        </div>
      </header>
      <section className="form-section">
        <h3 style={{ marginBottom: "6vh" }}>لطفا اطلاعات خودرو را وارد کنید</h3>
        <form style={{ marginBottom: "5vh" }}>
          <label>برند و تیپ</label>
          <Select className="select" placeholder="برند و تیپ" options={brand} />
          <label>سال تولید (بر اساس تیپ)</label>
          <Select
            className="select"
            placeholder="سال تولید (بر اساس تیپ)"
            options={year}
            id="year"
            // value={year_}
            onChange={(e) => {
              console.log(e.label);
              // let res = year.find((item) => item.label === e.label);
              // console.log(res);
              console.log(typeof e.label);
              let str = parseInt(e.label, 10);
              // setyear(parseInt(str, 10));
              console.log(typeof str);
              setyear(str);
            }}
          />
          <label>رنگ</label>
          <Select
            className="select"
            placeholder="رنگ"
            options={color}
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            id="color"
            // value={color_value}
            onChange={(e) => {
              console.log(e.name);
              let res = color.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setcolor_value(res.value);
            }}
          />
          <label>کارکرد</label>
          <br />
          <input
            className="select"
            placeholder="کارکرد"
            type="text"
            id="input"
            onChange={(e) => {
              // console.log(e.target.value);
              // setusage_km(parseInt(e.target.value, 10));
              let str = parseInt(e.target.value, 10);
              console.log(typeof str);
              setusage_km(str);
            }}
          />
          <br />
          <label>وضعیت بدنه</label>
          <Select
            className="select"
            placeholder="وضعیت بدنه"
            options={body}
            id="body"
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            onChange={(e) => {
              console.log(e.name);
              let res = body.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setbody_status_value(res.value);
            }}
          />
          <label>وضعیت موتور</label>
          <Select
            className="select"
            placeholder="وضعیت موتور"
            options={motor}
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            id="motor"
            onChange={(e) => {
              console.log(e.name);
              let res = motor.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setmotor_status_value(res.value);
            }}
          />
          <label>وضعیت شاسی‌های جلو</label>
          <Select
            className="select"
            placeholder="وضعیت شاسی ها"
            options={frontChassis}
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            id="frontCh"
            onChange={(e) => {
              console.log(e.name);
              let res = frontChassis.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setfront_chassis_status_value(res.value);
            }}
          />
          <label>وضعیت شاسی‌های عقب</label>
          <Select
            className="select"
            placeholder="وضعیت شاسی ها"
            options={rearChassis}
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            id="rearCh"
            onChange={(e) => {
              console.log(e.name);
              let res = rearChassis.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setrear_chassis_status_value(res.value);
            }}
          />
          <label>مهلت بیمهٔ شخص ثالث</label>
          <Select
            className="select"
            placeholder="مهلت بیمه شخص ثالث"
            options={bime}
            id="insuranc"
            onChange={(e) => {
              console.log(e.label);
              let res = bime.find((item) => item.label === e.label);
              console.log(res);
              console.log(res.value);
              setinsurance(res.value);
            }}
          />
          <label>گیربکس</label>
          <Select
            className="select"
            placeholder="گیربکس"
            options={gearbox}
            getOptionLabel={(option) => option.name}
            getOptionValue={(option) => option.value}
            id="gearbox"
            onChange={(e) => {
              console.log(e.name);
              let res = gearbox.find((item) => item.name === e.name);
              console.log(res);
              console.log(res.value);
              setgearbox(res.value);
            }}
          />
          <button
            className="submit"
            onClick={(e) => {
              submit(e);
            }}
            type="submit"
          >
            ثبت
          </button>
        </form>
        <div className="predict">
          <p style={{ display: show ? "block" : "none" }}>
            {" "}
            قیمت تخمین زده شده: {predict}
          </p>
        </div>
      </section>
    </div>
  );
}

export default Home;
