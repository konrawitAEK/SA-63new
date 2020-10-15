import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import { InputLabel, MenuItem, Select } from '@material-ui/core';
import Typography from '@material-ui/core/Typography';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';
import Link from '@material-ui/core/Link';

import { EntPhysician } from '../../api/models/EntPhysician';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntPosition } from '../../api/models/EntPosition';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'to PositionAssingment' };
  const api = new DefaultApi();

  //const [Positionassingment, setPositionassingment] = React.useState(initialUserState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [physicians, setPhysicians] = React.useState<EntPhysician[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [positions, setPositions] = React.useState<EntPosition[]>([]);

  const [position, SetPosition] = useState(Number);
  const [department, SetDepartment] = useState(Number);
  const [userid, setUser] = useState(Number);
  const [daystart, SetDayStart] = useState(String);

  useEffect(() => {
    const getPhysicians = async () => {
      const p = await api.listPhysician({ limit: 10, offset: 0 });
      setLoading(false);
      setPhysicians(p);
    };
    getPhysicians();

    const getDepartments = async () => {
      const d = await api.listDepartment({ limit: 10, offset: 0 });
      setLoading(false);
      setDepartments(d);
    };
    getDepartments();

    const getPositions = async () => {
      const po = await api.listPosition({ limit: 10, offset: 0 });
      setLoading(false);
      setPositions(po);
    };
    getPositions();

  }, [loading]);


  const handletimeChange = (event: any) => {
    SetDayStart(event.target.value as string);
  };

  const CreatePositionassingment = async () => {
    const positionassingment = {
      physicianid: userid,
      departmentid: department,
      positionid: position,
      dayStart: daystart + "T00:00:00Z"
    }

    const res: any = await api.createPositionassingment({ positionassingment: positionassingment });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };


  const physician_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const department_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetDepartment(event.target.value as number);
  };

  const position_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetPosition(event.target.value as number);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to PositionAssingment'}`}
        subtitle="Select the Position you want to be in."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
          <Link
            component="button"
            variant="body2"
            onClick={() => {
              console.info("I'm a button.");
            }}
          >
            Logout
    </Link>
        </Typography>

      </Header>
      <Content>

        <ContentHeader title="PositionAssingment">
          <Button variant="outlined" color="secondary" href="#outlined-buttons">
            LogOut
        </Button>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            

            <TableCell align="left">

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <TextField
                  id="DayStart"
                  label="DayStart"
                  type="date"
                  value={daystart}
                  onChange={handletimeChange}
                  //defaultValue="2020-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}

                />
              </FormControl>


              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="physician_id-label">Physician</InputLabel>
                <Select
                  labelId="physician_id-label"
                  label="Physician"
                  id="physician_id"
                  value={userid}
                  onChange={physician_id_handleChange}
                  style={{ width: 300 }}
                >
                  {physicians.map((item: EntPhysician) =>
                    <MenuItem value={item.id}>{item.eMAIL}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="department_id-label">Department</InputLabel>
                <Select
                  labelId="department_id-label"
                  label="Department"
                  id="department_id"
                  value={department}
                  onChange={department_id_handleChange}
                  style={{ width: 300 }}
                >
                  {departments.map((item: EntDepartment) =>
                    <MenuItem value={item.id}>{item.departmentname}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="position_id-label">Position</InputLabel>
                <Select
                  labelId="position_id-label"
                  label="Position"
                  id="position_id"
                  value={position}
                  onChange={position_id_handleChange}
                  style={{ width: 300 }}
                >
                  {positions.map((item: EntPosition) =>
                    <MenuItem value={item.id}>{item.nameposition}</MenuItem>)}
                </Select>
              </FormControl>
            </TableCell>
            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  onClick={() => {
                    CreatePositionassingment();
                  }}
                  variant="contained"
                  color="primary"
                  style={{ marginLeft: 545, width: 200 }}
                >
                  SAVE DATA
             </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/"
                  variant="contained"
                >
                  Back
             </Button>
              </TableCell>

            </div>
          </form>
        </div>
      </Content>

    </Page>
  );
}
