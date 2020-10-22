import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';


const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'to PositionAssingment' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
       subtitle="Select the department and position you want to be in."
     >
         
          <Button variant="outlined" color="secondary" href="#outlined-buttons" component={RouterLink} to="/">
            LogOut
        </Button>          
        
     </Header>
     <Content>
       <ContentHeader title="PositionAssingment">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add Position
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
export default WelcomePage;