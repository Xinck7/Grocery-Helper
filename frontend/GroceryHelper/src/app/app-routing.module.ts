import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IngredientsComponent } from './components/ingredients/ingredients.component';
import { ItemsComponent } from './components/items/items.component';
import { ListComponent } from './components/list/list.component';
import { RecipesComponent } from './components/recipes/recipes.component';
import { LoginComponent } from './components/login/login.component';
import { UserProfileComponent } from "./components/user-profile/user-profile.component";
import { AuthGuard } from "./guards/auth.guard";
import { RegisterComponent } from "./components/register/register.component"
import { NotFoundComponent } from "./components/not-found/not-found.component"


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
    path: '',
    component: LoginComponent 
  },
  {
    path: 'register',
    component: RegisterComponent
  },
  {
    path: '**',
    component: NotFoundComponent,
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
