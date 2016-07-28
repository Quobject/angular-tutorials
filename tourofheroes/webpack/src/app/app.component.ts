import { Component } from '@angular/core';
import { ROUTER_DIRECTIVES } from '@angular/router';

import { TasklistComponent } from './tasklist.component';
import { Task1Component } from './task1.component';
import { Task2Component } from './task2.component';

import '../../public/css/styles.css';

@Component({
  selector: 'my-app',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  directives: [ROUTER_DIRECTIVES],
  precompile: [TasklistComponent, Task1Component, Task2Component]
})
export class AppComponent { }
