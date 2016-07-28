import { Component } from '@angular/core';
import { ROUTER_DIRECTIVES } from '@angular/router';

import { Task1Component } from './task1.component';
import { Task2Component } from './task2.component';



@Component({
  selector: 'tasklist',
  templateUrl: './tasklist.component.html',
  styleUrls: ['./tasklist.component.css'],
  directives: [ROUTER_DIRECTIVES],
  precompile: [Task1Component, Task2Component]
})

export class TasklistComponent { }
