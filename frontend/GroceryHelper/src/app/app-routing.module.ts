import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IngredientsComponent } from './ingredients/ingredients.component';
import { ItemsComponent } from './items/items.component';
import { ListComponent } from './list/list.component';
import { RecipesComponent } from './recipes/recipes.component';

const routes: Routes = [
  {
    path: 'ingredients',
    component: IngredientsComponent
  },
  {
    path: 'items',
    component: ItemsComponent
  },
  {
    path: 'list',
    component: ListComponent
  },
  {
    path: 'recipes',
    component: RecipesComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
