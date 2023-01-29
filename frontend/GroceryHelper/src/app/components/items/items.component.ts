import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  // Create new item with form control for the item
  newItem(){
    
  }

  // Get exisiting items owned by current logged in user
  getItems(){

  }

  // Edit items
  editItem(){

  }
  
  // Get all items
  getAllItems(){

  }

  // add selected items to a list
  addSelectedToList(){

  }


}
