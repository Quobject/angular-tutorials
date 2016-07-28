import { provideRouter, RouterConfig } from '@angular/router';

import { TasklistComponent }  from './tasklist.component';
import { Task1Component }  from './task1.component';
import { Task2Component }    from './task2.component';

const routes: RouterConfig = [
  { path: 'tasklist', component: TasklistComponent },
  { path: 'task1', component: Task1Component },
  { path: 'task2', component: Task2Component },  { path:  '',  redirectTo: '/tasklist', pathMatch: 'full'  } 
];

export const appRouterProviders = [
  provideRouter(routes)
];
