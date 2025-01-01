import { useEffect } from 'react';
import { useProduct } from '@/hooks/useProduct';
import ProductList from '@/components/product/ProductList';
import { Loading } from '@/components/common/Loading';

const ListPage = () => {
  const { fetchProducts, loading } = useProduct();

  useEffect(() => {
    fetchProducts();
  }, [fetchProducts]);

  if (loading) return <Loading />;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-6">Products</h1>
      <ProductList />
    </div>
  );
};

export default ListPage;