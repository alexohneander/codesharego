import { BrowserRouter, Route, Switch } from "react-router-dom";

import Fiber from "./components/Fiber";
import React from "./components/React";
import NotFound from "./components/NotFound";

const App = () => (
  <BrowserRouter>
    <Switch>
      <Route path="/" component={Fiber} exact />
      <Route path="/room" component={React} />
      <Route path="*" component={NotFound} />
    </Switch>
  </BrowserRouter>
);

export default App;
