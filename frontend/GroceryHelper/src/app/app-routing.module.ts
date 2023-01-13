import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IngredientsComponent } from './ingredients/ingredients.component';
import { ItemsComponent } from './items/items.component';
import { ListComponent } from './list/list.component';
import { RecipesComponent } from './recipes/recipes.component';
import { LoginComponent } from './login/login.component';
import { UserProfileComponent } from "./user-profile/user-profile.component";
import { AuthGuard } from "./auth.guard";

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
  },
  {
    path: 'login',
    component: LoginComponent 
  },
  {
    path: 'user-profile/:id',
    component: UserProfileComponent, canActivate: [AuthGuard]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
