import { Component, OnInit } from '@angular/core';
import { takeLast } from 'rxjs';
import { Task } from '../model/task.model';


@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {
  checked = false;
  labelPosition: 'before' | 'after' = 'after';
  disabled = false;

  constructor() { }

  ngOnInit(): void {
  }
  public listLine : Task [] =[];

  addListLine(){
    this.listLine.push(new Task());
  }

  removeListLine(index:number){
    if(index > -1){
      this.listLine.splice(index,1);
    }
  }

}