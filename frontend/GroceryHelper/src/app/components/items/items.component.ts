import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { tap } from 'rxjs';
import { ItemsService } from 'src/app/services/items.service';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent implements OnInit {

  // itemForm!: FormGroup;
  itemList: any;
  constructor(
    // private formBuilder: FormBuilder
    private itemsService: ItemsService
  ) { }

  ngOnInit(): void {
  // this.itemForm = this.formBuilder.group({
  //   name: ['', Validators.required],
  //   location: ['', Validators.required],
  //   ailse: ['', Validators.required],
  //   price: ['', Validators.required],
  //   picture: ['', Validators.required],
  //   added_by: ['', Validators.required],
  //   show_only_your_added_by: ['', Validators.required]
  // });
    this.itemList = this.getItems();
  }

  // Create new item with form control for the item
  newItem(){
    
  }

  // Get exisiting items owned by current logged in user
  getItems(){
    this.itemsService.getItems().subscribe((data: any) => {
      this.itemList = data.results
    });
  }

  // Edit items
  editItem(){

  }

  // Get all items
  getAllItems(){

  }

  // add selected items to a list
  addSelectedToList(){
    console.log("you TOTALLY added a new item");
  }


}
